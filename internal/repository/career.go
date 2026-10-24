package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"career-arch-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CareerRepository implements CareerRepository interface
type CareerRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewCareerRepository creates a new repository instance
func NewCareerRepository(client *mongo.Client) models.CareerRepository {
	return &CareerRepository{
		client:     client,
		collection: client.Database("career_arch").Collection("careers"),
	}
}

// FindAll returns all careers
func (r *CareerRepository) FindAll() ([]models.Career, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to find careers: %w", err)
	}
	defer cursor.Close(ctx)

	var careers []models.Career
	if err := cursor.All(ctx, &careers); err != nil {
		return nil, fmt.Errorf("failed to decode careers: %w", err)
	}

	return careers, nil
}

// FindByID returns a career by ID
func (r *CareerRepository) FindByID(id string) (*models.Career, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var career models.Career
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %w", err)
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&career)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("career not found")
		}
		return nil, fmt.Errorf("failed to find career: %w", err)
	}

	return &career, nil
}

// FindBySlug returns a career by slug
func (r *CareerRepository) FindBySlug(slug string) (*models.Career, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var career models.Career
	err := r.collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&career)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("career not found")
		}
		return nil, fmt.Errorf("failed to find career: %w", err)
	}

	return &career, nil
}

// FindByCategory returns careers by category
func (r *CareerRepository) FindByCategory(category string) ([]models.Career, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"category": category})
	if err != nil {
		return nil, fmt.Errorf("failed to find careers: %w", err)
	}
	defer cursor.Close(ctx)

	var careers []models.Career
	if err := cursor.All(ctx, &careers); err != nil {
		return nil, fmt.Errorf("failed to decode careers: %w", err)
	}

	return careers, nil
}

// FindByTags returns careers matching any of the given tags
func (r *CareerRepository) FindByTags(tags []string) ([]models.Career, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"tags": bson.M{"$in": tags}})
	if err != nil {
		return nil, fmt.Errorf("failed to find careers: %w", err)
	}
	defer cursor.Close(ctx)

	var careers []models.Career
	if err := cursor.All(ctx, &careers); err != nil {
		return nil, fmt.Errorf("failed to decode careers: %w", err)
	}

	return careers, nil
}

// Create inserts a new career
func (r *CareerRepository) Create(career *models.Career) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	career.CreatedAt = time.Now().Format(time.RFC3339)
	career.UpdatedAt = career.CreatedAt

	_, err := r.collection.InsertOne(ctx, career)
	if err != nil {
		return fmt.Errorf("failed to create career: %w", err)
	}

	return nil
}

// Update modifies an existing career
func (r *CareerRepository) Update(career *models.Career) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	career.UpdatedAt = time.Now().Format(time.RFC3339)

	filter := bson.M{"_id": career.ID}
	update := bson.M{
		"$set": bson.M{
			"name":           career.Name,
			"slug":           career.Slug,
			"description":    career.Description,
			"category":       career.Category,
			"difficulty":     career.Difficulty,
			"tags":           career.Tags,
			"skills":         career.Skills,
			"content_blocks": career.ContentBlocks,
			"roadmap":        career.Roadmap,
			"metrics":        career.Metrics,
			"updated_at":     career.UpdatedAt,
		},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update career: %w", err)
	}

	return nil
}

// Delete removes a career by ID
func (r *CareerRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %w", err)
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete career: %w", err)
	}

	return nil
}

// Count returns total number of careers
func (r *CareerRepository) Count() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, fmt.Errorf("failed to count careers: %w", err)
	}

	return count, nil
}

// InitializeDatabase creates indexes and seeds initial data
func InitializeDatabase(client *mongo.Client) error {
	log.Println("Initializing database...")

	// Create indexes
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "slug", Value: 1}},
			Options: options.Index().
				SetUnique(true).
				SetName("slug_idx"),
		},
		{
			Keys: bson.D{{Key: "category", Value: 1}},
			Options: options.Index().
				SetName("category_idx"),
		},
		{
			Keys: bson.D{{Key: "tags", Value: 1}},
			Options: options.Index().
				SetName("tags_idx"),
		},
		{
			Keys: bson.D{{Key: "difficulty", Value: 1}},
			Options: options.Index().
				SetName("difficulty_idx"),
		},
	}

	_, err := client.Database("career_arch").Collection("careers").Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}

	log.Println("Database indexes created successfully")
	return nil
}
