package models

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Initialize Firestore client
func InitFirestore(ctx context.Context, keyFile string) (*firestore.Client, error) {
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(keyFile))
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Add a new Meal document
func AddMeal(ctx context.Context, client *firestore.Client, mid int, plannerID string, mealType string, dishID string) (string, error) {
	data := map[string]interface{}{
		"mid":     mid,
		"planner": plannerID,
		"meal":    mealType,
		"dish":    dishID,
	}
	docRef, _, err := client.Collection("meals").Add(ctx, data)
	if err != nil {
		return "", err
	}
	return docRef.ID, nil
}

// Add a new Food document
func AddFood(ctx context.Context, client *firestore.Client, fid int, name string, mealType string) (string, error) {
	data := map[string]interface{}{
		"fid":  fid,
		"name": name,
		"meal": mealType,
	}
	docRef, _, err := client.Collection("foods").Add(ctx, data)
	if err != nil {
		return "", err
	}
	return docRef.ID, nil
}

// Add a new Planner document
func AddPlanner(ctx context.Context, client *firestore.Client, pid int, userID int, planName string, createdAt string) (string, error) {
	data := map[string]interface{}{
		"pid":       pid,
		"user":      userID,
		"planname":  planName,
		"created_at": createdAt,
	}
	docRef, _, err := client.Collection("planners").Add(ctx, data)
	if err != nil {
		return "", err
	}
	return docRef.ID, nil
}

// Retrieve all documents from a collection
func RetrieveDocuments(ctx context.Context, client *firestore.Client, collection string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	iter := client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		results = append(results, doc.Data())
	}
	return results, nil
}
