const request = require('supertest')

const SERVER_ADDRESS  = "localhost:3000"
const PAYLOAD_GENERAL = {
    "firstName": "__test_FirstName",
    "lastName" : "__test_LastName",
}

describe("[GENERAL] LastDisco /user api endpoint", () => {
    describe("POST request", () => {
        it("Is available", async () => {
            await request(SERVER_ADDRESS)
                .post  ("/user")
                .send  (PAYLOAD_GENERAL)
                .expect(200)
        })

        it("Properly stores the recieved data", async () => {
            const reponse = await request(SERVER_ADDRESS)
                .post  ("/user")
                .send  (PAYLOAD_GENERAL)
            const responseJSON = JSON.parse(reponse.text)

            expect(responseJSON['FirstName']).toEqual(PAYLOAD_GENERAL['firstName'])
            expect(responseJSON['LastName']) .toEqual(PAYLOAD_GENERAL['lastName'])
        })
    })

    describe("GET request", () => {
        it("Is available", async () => {
            const response = await request(SERVER_ADDRESS)
                .get   ("/user")
                .expect(200) 
            const responseJSON = JSON.parse(response.text)
            const subject      = responseJSON[0]

            expect(subject['FirstName']).toEqual(PAYLOAD_GENERAL['firstName'])
            expect(subject['LastName']) .toEqual(PAYLOAD_GENERAL['lastName'])
        })
    })

    describe("PUT request", () => {

        it("Is not available", async () => {
            const response = await request(SERVER_ADDRESS)
                .put ("/user")
                .send(PAYLOAD_GENERAL)
            expect(response.statusCode).not.toBe(200)
        })
    })

    describe("DELETE request", () => {

        it("Is not available", async () => {
            const response = await request(SERVER_ADDRESS)
                .put ("/user")
                .send(PAYLOAD_GENERAL)
            expect(response.statusCode).not.toBe(200)
        })
    })
})

describe("[GENERAL] LastDisco /user/{id} api endpoint", () => {

    describe("GET request", () => {
        it("Is available", async () => {
            await request(SERVER_ADDRESS)
                .post("/user")
                .send(PAYLOAD_GENERAL)

            await request(SERVER_ADDRESS)
                .get   ("/user/1")
                .expect(200)
        })

        it("Responds with the correct data", async () => {
            const reponse = await request(SERVER_ADDRESS)
                .get ("/user/1")
                .send(PAYLOAD_GENERAL)
            const responseJSON = JSON.parse(reponse.text)

            expect(responseJSON['FirstName']).toEqual(PAYLOAD_GENERAL['firstName'])
            expect(responseJSON['LastName']) .toEqual(PAYLOAD_GENERAL['lastName'])
        })
    })

    describe("DELETE request", () => {

        it("Is available", async () => {
            await request(SERVER_ADDRESS)
                .delete("/user/1")
                .expect(200)
        })
    })

    describe("POST request", () => {

        it("Is not available", async () => {
            const response = await request(SERVER_ADDRESS)
                .post("/user/1")
                .send(PAYLOAD_GENERAL)
            expect(response.statusCode).not.toBe(200)
        })
    })
})

describe("[ACTION] CRUD the user", () => {
    var bakedProductID = undefined

    it("Is creating the user", async () => {
        const response = await request(SERVER_ADDRESS)
            .post  ("/user")
            .send  (PAYLOAD_GENERAL)
            .expect(200)

        bakedProductID = JSON.parse(response.text)["ID"]
    })

    it("Is reading the product", async () => {
        const response = await request(SERVER_ADDRESS)
            .get("/user/2")
        responseJSON = JSON.parse(response.text)

        expect(responseJSON['FirstName']).toEqual(PAYLOAD_GENERAL['firstName'])
        expect(responseJSON['LastName']) .toEqual(PAYLOAD_GENERAL['lastName'])
    })

    it("Is updating the product", async () => {
        payloadCustom              = PAYLOAD_GENERAL
        payloadCustom["firstName"] = "__TestUserUpdated"

        var response = await request(SERVER_ADDRESS)
            .put("/user/2")
            .send(payloadCustom)
        expect(response.status).toBe(200)

        response = await request(SERVER_ADDRESS)
            .get("/user/2")
        responseJSON = JSON.parse(response.text)

        expect(responseJSON["FirstName"]).toEqual(payloadCustom["firstName"])
    })
})