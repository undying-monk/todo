from sklearn import tree

# 1. Features: [Weight, Texture (1=Smooth, 0=Bumpy)]
# These are our "clues"
features = [[140, 1], [130, 1], [150, 0], [170, 0]]

# 2. Labels: [0 = Apple, 1 = Orange]
# These are the correct "answers"
labels = [0, 0, 1, 1]

# 3. Initialize the Model (a Decision Tree)
# Think of this as an empty brain
model = tree.DecisionTreeClassifier()

# 4. Train the Model (The "Learning" phase)
# We show the model the clues and the answers
model = model.fit(features, labels)

# 5. Predict (The "Application" phase)
# We give it a new fruit: 160g and Bumpy (0)
prediction = model.predict([[160, 0]])

print(prediction)