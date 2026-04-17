import tensorflow as tf
from tensorflow.keras import layers

# 1. Define the Architecture (The "Layers")
model = tf.keras.Sequential([
    # Hidden Layer: 10 "neurons" looking for patterns
    layers.Dense(10, activation='relu', input_shape=(2,)), 
    
    # Output Layer: 1 neuron giving a final 0 or 1 result
    layers.Dense(1, activation='sigmoid')                 
])

# 2. Compile (The "Learning Strategy")
model.compile(optimizer='adam', 
              loss='binary_crossentropy', 
              metrics=['accuracy'])

# 3. Fit (The "Training")
# model.fit(features, labels, epochs=50)