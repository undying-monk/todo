package classify

import (
	"os"
	"testing"
)

func TestClassify(t *testing.T) {
	// Classify panics on error, so we can't easily test it without a model.bin
	// but we can test if it doesn't panic if everything is set up.
}

func TestGetPredictions_WithEnv(t *testing.T) {
	os.Setenv("FASTTEXT_MODEL_PATH", "non_existent_model.bin")
	defer os.Unsetenv("FASTTEXT_MODEL_PATH")

	_, err := GetPredictions(os.Getenv("FASTTEXT_MODEL_PATH"), "test text")
	if err == nil {
		t.Fatal("Expected error when model path from env is missing, got nil")
	}
}

func TestGetPredictions_ModelNotFound(t *testing.T) {
	_, err := GetPredictions("missing_model.bin", "test text")
	if err == nil {
		t.Fatal("Expected error when model is missing, got nil")
	}
}

