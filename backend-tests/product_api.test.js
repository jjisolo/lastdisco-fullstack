const request = require('supertest')

const SERVER_ADDRESS  = "localhost:3000"
const PAYLOAD_GENERAL = {
    "shortName"  : "__test_Product",
    "description": "__test_Description",
    "count"      : 100,
    "price"      : "12.99"
}

describe("[GENERAL] LastDisco /product api endpoint", () => {
    describe("POST request", () => {
        it("Is available", async () => {
            await request(SERVER_ADDRESS)
                .post  ("/product")
                .send  (PAYLOAD_GENERAL)
                .expect(200)
        })

        it("Properly stores the recieved data", async () => {
            const reponse = await request(SERVER_ADDRESS)
                .post  ("/product")
                .send  (PAYLOAD_GENERAL)
            const responseJSON = JSON.parse(reponse.text)

            expect(responseJSON['ShortName'])  .toEqual(PAYLOAD_GENERAL['shortName'])
            expect(responseJSON['Description']).toEqual(PAYLOAD_GENERAL['description'])
            expect(responseJSON['Count'])      .toEqual(PAYLOAD_GENERAL['count'])
            expect(responseJSON['Price'])      .toEqual(PAYLOAD_GENERAL['price'])
        })
    })

    describe("GET request", () => {

        it("Is available", async () => {
            const response = await request(SERVER_ADDRESS)
                .get   ("/product")
                .expect(200) 
            const responseJSON = JSON.parse(response.text)
            const subject      = responseJSON[0]

            expect(subject['ShortName'])  .toEqual(PAYLOAD_GENERAL['shortName'])
            expect(subject['Description']).toEqual(PAYLOAD_GENERAL['description'])
            expect(subject['Count'])      .toEqual(PAYLOAD_GENERAL['count'])
            expect(subject['Price'])      .toEqual(PAYLOAD_GENERAL['price'])
        })
    })

    describe("PUT request", () => {

        it("Is not available", async () => {
            const response = await request(SERVER_ADDRESS)
                .put ("/product")
                .send(PAYLOAD_GENERAL)
            expect(response.statusCode).not.toBe(200)
        })
    })

    describe("DELETE request", () => {

        it("Is not available", async () => {
            const response = await request(SERVER_ADDRESS)
                .put ("/product")
                .send(PAYLOAD_GENERAL)
            expect(response.statusCode).not.toBe(200)
        })
    })
})

describe("[GENERAL] LastDisco /product/{id} api endpoint", () => {

    describe("GET request", () => {
        it("Is available", async () => {
            await request(SERVER_ADDRESS)
                .post("/product")
                .send(PAYLOAD_GENERAL)

            await request(SERVER_ADDRESS)
                .get   ("/product/1")
                .expect(200)
        })

        it("Responds with the correct data", async () => {
            const reponse = await request(SERVER_ADDRESS)
                .get ("/product/1")
                .send(PAYLOAD_GENERAL)
            const responseJSON = JSON.parse(reponse.text)

            expect(responseJSON['ShortName'])  .toEqual(PAYLOAD_GENERAL['shortName'])
            expect(responseJSON['Description']).toEqual(PAYLOAD_GENERAL['description'])
            expect(responseJSON['Count'])      .toEqual(PAYLOAD_GENERAL['count'])
            expect(responseJSON['Price'])      .toEqual(PAYLOAD_GENERAL['price'])
        })
    })

    describe("DELETE request", () => {

        it("Is available", async () => {
            await request(SERVER_ADDRESS)
                .delete("/product/1")
                .expect(200)
        })
    })

    describe("POST request", () => {

        it("Is not available", async () => {
            const response = await request(SERVER_ADDRESS)
                .post("/product/1")
                .send(PAYLOAD_GENERAL)
            expect(response.statusCode).not.toBe(200)
        })
    })
})

describe("[ACTION] CRUD the product", () => {
    var bakedProductID = undefined

    it("Is creating the product", async () => {
        const response = await request(SERVER_ADDRESS)
            .post  ("/product")
            .send  (PAYLOAD_GENERAL)
            .expect(200)

        bakedProductID = JSON.parse(response.text)["ID"]
    })

    it("Is reading the product", async () => {
        const response = await request(SERVER_ADDRESS)
            .get("/product/2")
        responseJSON = JSON.parse(response.text)

        expect(responseJSON['ShortName'])  .toEqual(PAYLOAD_GENERAL['shortName'])
        expect(responseJSON['Description']).toEqual(PAYLOAD_GENERAL['description'])
        expect(responseJSON['Count'])      .toEqual(PAYLOAD_GENERAL['count'])
        expect(responseJSON['Price'])      .toEqual(PAYLOAD_GENERAL['price'])
    })

    it("Is updating the product", async () => {
        payloadCustom              = PAYLOAD_GENERAL
        payloadCustom["shortName"] = "__TestProductUpdated"

        var response = await request(SERVER_ADDRESS)
            .put("/product/2")
            .send(payloadCustom)
        expect(response.status).toBe(200)

        response = await request(SERVER_ADDRESS)
            .get("/product/2")
        responseJSON = JSON.parse(response.text)

        expect(responseJSON["ShortName"]).toEqual(payloadCustom["shortName"])
    })
})