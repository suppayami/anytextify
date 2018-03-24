package services_test

import (
	"testing"

	"github.com/suppayami/anytextify/services"
)

func TestSimpleTransform(t *testing.T) {
	nyanify := services.Nyanify{}
	if nyanify.Transform("test") != "nyan" {
		t.Error(`nyanify.Transform("test") != "nyan"`)
	}
}

func TestTransformWithPunctuation(t *testing.T) {
	nyanify := services.Nyanify{}
	if nyanify.Transform("test!abc") != "nyan!nya" {
		t.Error(`nyanify.Transform("test!abc") != "nyan!nya"`)
	}
}

func TestTransformWithManyWords(t *testing.T) {
	nyanify := services.Nyanify{}
	if nyanify.Transform("test!abc loli moe") != "nyan!nya nyan nya" {
		t.Error(`nyanify.Transform("test!abc loli moe") != "nyan!nya nyan nya"`)
	}
}
