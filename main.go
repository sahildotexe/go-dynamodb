package main

import (
	"fmt"

	"github.com/sahildotexe/go-dynamodb/infrastructure"
	"github.com/sahildotexe/go-dynamodb/initializers"
	"github.com/sahildotexe/go-dynamodb/services"
)

func main() {
	initializers.LoadEnvVariables()
	// load AWS config and dynamodb client
	config := infrastructure.NewAWSConfig()
	client := infrastructure.NewDynamoDBClient(config)

	// CREATE CHAT TABLE
	// chatTable, err := services.CreateChatTable(client)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf(
	// 	"Table ID: %s \nTable Name: %s\n\n",
	// 	*chatTable.TableDescription.TableId,
	// 	*chatTable.TableDescription.TableName,
	// )

	// add chat to db
	// ChatID := uuid.New()
	// chatData := services.ChatDataType{
	// 	UserID:    "dbebf8e1-a375-4f9b-af6d-41f057e7b49b",
	// 	ChatID:    ChatID.String(),
	// 	Title:     "Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts.",
	// 	CreatedAt: time.Now().Unix(),
	// }

	var err error

	// _, err = services.Create(client, chatData)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("New Chat Created.")

	// GET CHAT
	// chats, err := services.GetUserChats(client, "dbebf8e1-a375-4f9b-af6d-41f057e7b49b")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// loop over the chats
	// for _, chat := range *chats {
	// 	fmt.Printf(
	// 		"User ID: %s \nChat ID: %s \nChat Title: %s\n\n",
	// 		chat.UserID,
	// 		chat.ChatID,
	// 		chat.Title,
	// 	)
	// }

	// get single chat item
	// chat, err := services.GetSingleChat(
	// 	client,
	// 	"dbebf8e1-a375-4f9b-af6d-41f057e7b49b", //user_id
	// 	"9cab3da2-0bc4-40d6-a634-842c7a1bb868", //chat_id
	// )

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf(
	// 	"\n\nSingle Chat \nUser ID: %s \nChat ID: %s \nChat Title: %s\n\n",
	// 	chat.UserID,
	// 	chat.ChatID,
	// 	chat.Title,
	// )

	// UPDATE CHAT
	// updatedFields, err := services.UpdateChat(client, services.ChatDataType{
	// 	UserID: "dbebf8e1-a375-4f9b-af6d-41f057e7b49b",
	// 	ChatID: "9cab3da2-0bc4-40d6-a634-842c7a1bb868",
	// 	Title:  "The far away mountain loomed in the distance, its peak shrouded in mist and mystery.",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("Updated %s \n\n", updatedFields.Attributes["chat_id"])

	// delete chat
	err = services.DeleteChat(
		client,
		"dbebf8e1-a375-4f9b-af6d-41f057e7b49b",
		"9cab3da2-0bc4-40d6-a634-842c7a1bb868",
	)

	if err != nil {
		fmt.Println(err)
	}

}
