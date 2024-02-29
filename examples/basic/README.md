# Simple bot implementation

Here you have an very simple example to create a simple bot to handle lobby creation/lunch  step by step with **go-dota2** library.

## Settings

- You must set your credentials and info in a **.env** file, check **.env.example**.

## Run the project

- ``` go run main.go```

## Test your bot

It's very easy, just use some API platform like Postman, Insomia, RapidAPI or anyone that you prefer.

Now you can start 
1. Send a **post** petition into:  localhost:8080/lobby 
    - Verifies in the dota 2 client that the lobby has been created satisfactorily.
2. Now you can send invitations to any of your friends: Send a **post** petition into: localhost:8080/invite/:steamid
    - Of course, set the param correctly.
3. Whenever you want, you can lunch the lobby: Send a **post** petition into: localhost:8080/start
4. Finally you can leave the lobby so your bot should be ready to keep working creating more lobbies: Send a **post**  petition into: localhost:8080/leave

> [!IMPORTANT]
> This is a simple guide to implement this library, to do advanced things you should check the entire docs: 
https://pkg.go.dev/github.com/paralin/go-dota2#section-readme




