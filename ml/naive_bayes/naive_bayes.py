import numpy as np
import matplotlib.pyplot as plt
import nltk
from nltk.corpus import stopwords
from nltk import word_tokenize
import w1_unittest
import pandas as pd
import string

# nltk.download('stopwords')
# nltk.download('punkt_tab')

def preprocess_emails(df):
    # shuffle
    df = df.sample(frac=1, ignore_index = True, random_state=42)

    # Remove Subject: on text column, convert to np
    X = df.text.apply(lambda x: x[9:]).to_numpy()
    Y = df.spam.to_numpy()
    return X,Y

# text,spam (0,1)
dataframe_emails = pd.read_csv('./data/emails.csv')
dataframe_emails.head()

print(f"Number of emails: {len(dataframe_emails)}")
print(f"Proportion of spam emails: {dataframe_emails.spam.sum()/len(dataframe_emails):.4f}")
print(f"Proportion of ham emails: {1-dataframe_emails.spam.sum()/len(dataframe_emails):.4f}")

X, Y = preprocess_emails(dataframe_emails)
print(X[:5])
print(Y[:5])

email_index = 30
print(f"Email index {email_index}: {X[email_index]}\n\n")
print(f"Class: {Y[email_index]}")

def preprocess_text(X):
    stop = set(stopwords.words('english') + list(string.punctuation))
    if isinstance(X, str):
        X = np.array([X])
    X_preprocessed = []
    for _,email in enumerate(X):
        email = np.array([i.lower() for i in word_tokenize(email) if i.lower() not in stop]).astype(X.dtype)
        X_preprocessed.append(email)
    if len(X) == 1:
        return X_preprocessed[0]
    return X_preprocessed

def get_word_frequency(X, Y):
    num_emails = len(X)
    frequency_words = {}

    for i in range (num_emails):
        email = X[i]
        email = set(email) # remove duplicate words
        for word in email:
            if word not in frequency_words:
                frequency_words[word] = {"spam": 1, "ham":1}
            key = ""
            if Y[i] == 1:
                key = "spam"
            else:
                key = "ham"
            frequency_words[word][key]+=1
    return frequency_words

X_treated = preprocess_text(X)
print(X_treated)

email_index = 989
print(f"Email before preprocessing: {X[email_index]}")
print(f"Email after preprocessing: {X_treated[email_index]}")


TRAIN_SIZE = int(0.80*len(X_treated)) # 80% of the samples will be used to train.

X_train = X_treated[:TRAIN_SIZE]
Y_train = Y[:TRAIN_SIZE]
X_test = X_treated[TRAIN_SIZE:]
Y_test = Y[TRAIN_SIZE:]

print(f"Proportion of spam in train dataset: {sum(Y_train == 1)/len(Y_train):.4f}")
print(f"Proportion of spam in test dataset: {sum(Y_test == 1)/len(Y_test):.4f}")

test_output = get_word_frequency([['like','going','river'], ['love', 'deep', 'river'], ['hate','river']], [1,0,0])
print(test_output)

w1_unittest.test_get_word_frequency(get_word_frequency)

word_frequency = get_word_frequency(X_train,Y_train)
print(word_frequency)

# To count the spam and ham emails, you may just sum the respective 1 and 0 values in the training dataset, since the convention is spam = 1 and ham = 0.
class_frequency = {'ham': sum(Y_train == 0), 'spam': sum(Y_train == 1)}
print(class_frequency)

# The idea is to compute  proportion of spam emails
proportion_spam = class_frequency['spam']/(class_frequency['ham'] + class_frequency['spam'])
print(f"The proportion of spam emails in training is: {proportion_spam:.4f}")

# P(word1 | spam) = P(word1 and spam email) / P(spam email)
# P(word1 | ham) = P(word1 and ham email) / P(ham email)
def prob_word_given_class(word, cls, word_frequency, class_frequency):
    if word not in word_frequency: # for skipping
        return 1
    amount_word_and_class = word_frequency[word][cls]
    p_word_given_class = amount_word_and_class/class_frequency[cls]
    return p_word_given_class

print(f"P(lottery | spam) = {prob_word_given_class('lottery', cls = 'spam', word_frequency = word_frequency, class_frequency = class_frequency)}")
print(f"P(lottery | ham) = {prob_word_given_class('lottery', cls = 'ham', word_frequency = word_frequency, class_frequency = class_frequency)}")
print(f"P(schedule | spam) = {prob_word_given_class('schedule', cls = 'spam', word_frequency = word_frequency, class_frequency = class_frequency)}")
print(f"P(schedule | ham) = {prob_word_given_class('schedule', cls = 'ham', word_frequency = word_frequency, class_frequency = class_frequency)}")

w1_unittest.test_prob_word_given_class(prob_word_given_class, word_frequency, class_frequency)

def prob_email_given_class(treated_email, cls, word_frequency, class_frequency):
    p = 1
    for word in treated_email:
        p *= prob_word_given_class(word, cls, word_frequency, class_frequency)
    return p


example_email = "Click here to win a lottery ticket and claim your prize!"
treated_email = preprocess_text(example_email)
prob_spam = prob_email_given_class(treated_email, cls = 'spam', word_frequency = word_frequency, class_frequency = class_frequency)
prob_ham = prob_email_given_class(treated_email, cls = 'ham', word_frequency = word_frequency, class_frequency = class_frequency)
print(f"Email: {example_email}\nEmail after preprocessing: {treated_email}\nP(email | spam) = {prob_spam}\nP(email | ham) = {prob_ham}")

w1_unittest.test_prob_email_given_class(prob_email_given_class, word_frequency, class_frequency)

def naive_bayes(treated_email, word_frequency, class_frequency, return_likelihood = False):
    p_spam = class_frequency["spam"] / (class_frequency["spam"] + class_frequency["ham"])
    p_ham = class_frequency["ham"] / (class_frequency["spam"] + class_frequency["ham"])
    prob_email_given_spam = prob_email_given_class(treated_email, "spam", word_frequency, class_frequency)
    prob_email_given_ham = prob_email_given_class(treated_email, "ham", word_frequency, class_frequency)

    spam_likelihood = p_spam * prob_email_given_spam
    han_likelihood = p_ham * prob_email_given_ham

    if return_likelihood:
        return spam_likelihood, han_likelihood

    if spam_likelihood >= han_likelihood:
        return 1
    return 0

example_email = "Click here to win a lottery ticket and claim your prize!"
treated_email = preprocess_text(example_email)

print(f"Email: {example_email}\nEmail after preprocessing: {treated_email}\nNaive Bayes predicts this email as: {naive_bayes(treated_email, word_frequency, class_frequency)}")

print("\n\n")
example_email = "Our meeting will happen in the main office. Please be there in time."
treated_email = preprocess_text(example_email)

print(f"Email: {example_email}\nEmail after preprocessing: {treated_email}\nNaive Bayes predicts this email as: {naive_bayes(treated_email, word_frequency, class_frequency)}")

w1_unittest.test_naive_bayes(naive_bayes, word_frequency, class_frequency)

def get_true_positive_negatives(Y_true, Y_pred, true_positive_negative):
    n = len(Y_true)
    count = 0
    if len(Y_true) != len(Y_pred):
        return "Number of true labels and predict labels must match!"
    for i in range(n):
        if Y_true[i] == Y_pred[i] == true_positive_negative:
            count +=1
    return count

def get_false_positives(Y_true, Y_pred):
    # Both Y_true and Y_pred must match in length.
    if len(Y_true) != len(Y_pred):
        return "Number of true labels and predict labels must match!"
    n = len(Y_true)

    false_positives = 0
    # Iterate over the number of elements in the list
    for i in range(n):
        # Get the true label for the considered email
        true_label_i = Y_true[i]
        # Get the predicted (model output) for the considered email
        predicted_label_i = Y_pred[i]
        # Increase the counter by 1 only if true_label_i = 0 and predicted_label_i = 0 (false positive)
        if true_label_i == 0 and predicted_label_i == 1:
            false_positives += 1
    return false_positives

# Create an empty list to store the predictions
Y_pred = []

# Iterate over every email in the test set
for email in X_test:
    # Perform prediction
    prediction = naive_bayes(email, word_frequency, class_frequency)
    # Add it to the list 
    Y_pred.append(prediction)

print(f"Y_test and Y_pred matches in length? Answer: {len(Y_pred) == len(Y_test)}")

# Get the number of true positives:
true_positives = get_true_positive_negatives(Y_test, Y_pred, 1)

# Get the number of true negatives:
true_negatives = get_true_positive_negatives(Y_test, Y_pred, 0)

print(f"The number of true positives is: {true_positives}\nThe number of true negatives is: {true_negatives}")

# Compute the accuracy by summing true negatives with true positives and dividing it by the total number of elements in the dataset. 
# Since both Y_pred and Y_test have the same length, it does not matter which one you use.
accuracy = (true_positives + true_negatives)/len(Y_test)

print(f"Accuracy is: {accuracy:.4f}")


email = "Please meet me in 2 hours in the main building. I have an important task for you."
# email = "You win a lottery prize! Congratulations! Click here to claim it"

# Preprocess the email
treated_email = preprocess_text(email)
# Get the prediction, in order to print it nicely, if the output is 1 then the prediction will be written as "spam" otherwise "ham".
prediction = "spam" if naive_bayes(treated_email, word_frequency, class_frequency) == 1 else "ham"
print(f"The email is: {email}\nThe model predicts it as {prediction}.")

def log_prob_email_given_class(treated_email, cls, word_frequency, class_frequency):
    # prob starts at 0 because it will be updated by summing it with the current log(P(word | class)) in every iteration
    prob = 0

    for word in treated_email: 
        # Only perform the computation for words that exist in the word frequency dictionary
        if word in word_frequency.keys(): 
            # Update the prob by summing it with log(P(word | class))
            prob += np.log(prob_word_given_class(word, cls,word_frequency, class_frequency))

    return prob

# Consider an email with only one word, so it reduces to compute the value P(word | class) or log(P(word | class)).
one_word_email = ['schedule']
word = one_word_email[0]
prob_spam = prob_email_given_class(one_word_email, cls = 'spam',word_frequency = word_frequency, class_frequency = class_frequency)
log_prob_spam = log_prob_email_given_class(one_word_email, cls = 'spam',word_frequency = word_frequency, class_frequency = class_frequency)
print(f"For word {word}:\n\tP({word} | spam) = {prob_spam}\n\tlog(P({word} | spam)) = {log_prob_spam}")

def log_naive_bayes(treated_email, word_frequency, class_frequency, return_likelihood = False):    
    # Compute P(email | spam) with the new log function
    log_prob_email_given_spam = log_prob_email_given_class(treated_email, cls = 'spam',word_frequency = word_frequency, class_frequency = class_frequency) 

    # Compute P(email | ham) with the function you defined just above
    log_prob_email_given_ham = log_prob_email_given_class(treated_email, cls = 'ham',word_frequency = word_frequency, class_frequency = class_frequency) 

    # Compute P(spam) using the class_frequency dictionary and using the formula #spam emails / #total emails
    p_spam = class_frequency['spam']/(class_frequency['ham'] + class_frequency['spam']) 

    # Compute P(ham) using the class_frequency dictionary and using the formula #ham emails / #total emails
    p_ham = class_frequency['ham']/(class_frequency['ham'] + class_frequency['spam']) 

    # Compute the quantity log(P(spam)) + log(P(email | spam)), let's call it log_spam_likelihood
    log_spam_likelihood = np.log(p_spam) + log_prob_email_given_spam 

    # Compute the quantity P(ham) * P(email | ham), let's call it ham_likelihood
    log_ham_likelihood = np.log(p_ham) + log_prob_email_given_ham 

    # In case of passing return_likelihood = True, then return the desired tuple
    if return_likelihood == True:
        return (log_spam_likelihood, log_ham_likelihood)
    
    # Compares both values and choose the class corresponding to the higher value. 
    # As the logarithm is an increasing function, the class with the higher value retains this property.
    if log_spam_likelihood >= log_ham_likelihood:
        return 1
    else:
        return 0
    
example_index = 4798
example_email = X[example_index]
treated_email = preprocess_text(example_email)
print(f"The email is:\n\t{example_email}\n\nAfter preprocessing:\n\t:{treated_email}")

spam_likelihood, ham_likelihood = naive_bayes(treated_email, word_frequency = word_frequency, class_frequency = class_frequency, return_likelihood = True)
print(f"spam_likelihood: {spam_likelihood}\nham_likelihood: {ham_likelihood}")

print(f"The example email is labeled as: {Y[example_index]}")
print(f"Naive bayes model classifies it as: {naive_bayes(treated_email, word_frequency, class_frequency)}")


log_spam_likelihood, log_ham_likelihood = log_naive_bayes(treated_email,word_frequency = word_frequency, class_frequency = class_frequency,return_likelihood = True)
print(f"log_spam_likelihood: {log_spam_likelihood}\nlog_ham_likelihood: {log_ham_likelihood}")

print(f"The example email is labeled as: {Y[example_index]}")
print(f"Log Naive bayes model classifies it as: {log_naive_bayes(treated_email,word_frequency = word_frequency, class_frequency = class_frequency)}")

# Let's get the predictions for the test set:

# Create an empty list to store the predictions
Y_pred = []


# Iterate over every email in the test set
for email in X_test:
    # Perform prediction
    prediction = log_naive_bayes(email,word_frequency = word_frequency, class_frequency = class_frequency)
    # Add it to the list 
    Y_pred.append(prediction)

# Get the number of true positives:
true_positives = get_true_positive_negatives(Y_test, Y_pred, 1)

# Get the number of true negatives:
true_negatives = get_true_positive_negatives(Y_test, Y_pred, 0)

print(f"The number of true positives is: {true_positives}\nThe number of true negatives is: {true_negatives}")

# Compute the accuracy by summing true negatives with true positives and dividing it by the total number of elements in the dataset. 
# Since both Y_pred and Y_test have the same length, it does not matter which one you use.
accuracy = (true_positives + true_negatives)/len(Y_test)

print(f"The accuracy is: {accuracy:.4f}")

def get_recall(Y_true, Y_pred):
    # Get the total number of spam emails. Since they are 1 in the data, it suffices summing all the values in the array Y.
    total_number_spams = Y_test.sum()
    # Get the true positives
    true_positives = get_true_positive_negatives(Y_true, Y_pred, 1)
    
    # Compute the recall
    recall = true_positives/total_number_spams
    return recall

# Use the Naive Bayes model (standard and log versions) to classify every email in the test dataset
Y_pred_naive_bayes = []
Y_pred_log_naive_bayes = []

for email in X_test:
 prediction = naive_bayes(email,word_frequency = word_frequency, class_frequency = class_frequency)
 log_prediction = log_naive_bayes(email,word_frequency = word_frequency, class_frequency = class_frequency)
 Y_pred_naive_bayes.append(prediction)
 Y_pred_log_naive_bayes.append(log_prediction)

# Compute the recall for both models
recall_naive_bayes = get_recall(Y_test, Y_pred_naive_bayes)
recall_log_naive_bayes = get_recall(Y_test, Y_pred_log_naive_bayes)
print(f"The proportion of spam emails the standard Naive Bayes model can correctly classify as spam (recall) is: {recall_naive_bayes:.4f}")
print(f"The proportion of spam emails the log Naive Bayes model can correctly classify as spam (recall) is: {recall_log_naive_bayes:.4f}")

def get_false_positives(Y_true, Y_pred):
    """
    Calculate the number of false positives instances in binary classification.

    Parameters:
    - Y_true (list): List of true labels (0 or 1) for each instance.
    - Y_pred (list): List of predicted labels (0 or 1) for each instance.

    Returns:
    - int: Number of false positives, where true label is 0 and predicted label is 1.
    """
    
    # Both Y_true and Y_pred must match in length.
    if len(Y_true) != len(Y_pred):
        return "Number of true labels and predict labels must match!"
    n = len(Y_true)

    false_positives = 0
    # Iterate over the number of elements in the list
    for i in range(n):
        # Get the true label for the considered email
        true_label_i = Y_true[i]
        # Get the predicted (model output) for the considered email
        predicted_label_i = Y_pred[i]
        # Increase the counter by 1 only if true_label_i = 0 and predicted_label_i = 0 (false positive)
        if true_label_i == 0 and predicted_label_i == 1:
            false_positives += 1
    return false_positives

# Count the ham emails mistakenly labeled as spam (false positives). Let's use the function get_false_positives you've seen above
 
false_positives_naive_bayes = get_false_positives(Y_test, Y_pred_naive_bayes)
false_positives_log_naive_bayes = get_false_positives(Y_test, Y_pred_log_naive_bayes)

print(f"Number of false positives in the standard Naive Bayes model: {false_positives_naive_bayes}")
print(f"Number of false positives in the log Naive Bayes model: {false_positives_log_naive_bayes}")

def get_precision(Y_true, Y_pred):
    # Get the true positives
    true_positives = get_true_positive_negatives(Y_true, Y_pred, 1)
    false_positives = get_false_positives(Y_true, Y_pred)
    precision = true_positives/(true_positives + false_positives)
    return precision

print(f"Precision of the standard Naive Bayes model: {get_precision(Y_test, Y_pred_naive_bayes):.4f}")
print(f"Precision of the log Naive Bayes model: {get_precision(Y_test, Y_pred_log_naive_bayes):.4f}")