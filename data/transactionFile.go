package data

import (
	"context"
	"log"
	"time"

	"github.com/BreCkver/st-codeChallenge/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	baseName       = "transaction"
	collectionName = "transactionFile"
)

/*TransacionFileSave */
func TransacionFileSave(txsFile *models.TransactionsFile) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientDB := Conexion()

	db := clientDB.Database(baseName)
	col := db.Collection(collectionName)
	result, err := col.InsertOne(ctx, txsFile)
	if err != nil {
		log.Printf("Error guardando info %v", err.Error())
		return "", err
	}

	objIdentifier, _ := result.InsertedID.(primitive.ObjectID)
	return objIdentifier.Hex(), nil
}

/*GetTransactionFile */
func GetTransactionFile(id string) (models.TransactionsFile, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	clientDB := Conexion()

	db := clientDB.Database(baseName)
	col := db.Collection(collectionName)

	var txFile models.TransactionsFile
	objId, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": objId}
	err := col.FindOne(ctx, condition).Decode(&txFile)

	if err != nil {
		return txFile, err
	}

	return txFile, err

}
