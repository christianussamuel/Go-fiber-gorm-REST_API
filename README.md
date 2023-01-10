# Go-fiber-gorm-REST_API
REST_API user for cv application

//UserProfile
Post("/api/profile", controllers.CreateUser)
Get("/api/profile/:profileCode", controllers.GetAUser)
Put("api/profile/:profileCode", controllers.EditAUser)

//workingexp
Get("/api/working-experience/:profileCode")
Put("/api/working-experience/:profileCode")

//employment
Post("/api/employment/:profileCode")


https://3f86-112-78-153-28.ap.ngrok.io/route

ex:
https://3f86-112-78-153-28.ap.ngrok.io/api/profile/1673388289