
context("Chef Endpoints", ()=>{
    describe("create endpoint", () => {
        it('should return 400 error if no name is provided', () => {
            const postBody ={
                "email": "joshmartin5001@gmail.com"
            }
            cy.request({
                method: "POST",
                url: "/chef/create",
                failOnStatusCode: false,
                body: JSON.stringify(postBody)
            }).then((response) =>{
                expect(response.status).to.eq(400)
            })
        })
        it('should return 400 error if no email is provided', () => {
            const postBody ={
                "name": "Josh Martin"
            }
            cy.request({
                method: "POST",
                url: "/chef/create",
                failOnStatusCode: false,
                body: JSON.stringify(postBody)
            }).then((response) =>{
                expect(response.status).to.eq(400)
            })
        })

        it('should return 400 error if no email is not valid', () => {
            const postBody ={
                "name": "Josh Martin",
                "email": "joshmartgmail.com"
            }
            cy.request({
                method: "POST",
                url: "/chef/create",
                failOnStatusCode: false,
                body: JSON.stringify(postBody)
            }).then((response) =>{
                expect(response.status).to.eq(400)
            })
        })

        it('should return status 200 if provide valid email and name', () => {
            const postBody ={
                "name": "Josh Martin",
                "email": "joshmartin@gmail.com"
            }
            cy.request({
                method: "POST",
                url: "/chef/create",
                failOnStatusCode: false,
                body: JSON.stringify(postBody)
            }).then((response) =>{
                expect(response.status).to.eq(200)
            })
        })
    })

    describe("/:id endpoints", () =>{

        it("should return status 400 if id does not exist", ()=>{
            const ID = "60d7f6d9c365af66a9f5915f"
            cy.request({
                method: "GET",
                url: `/chef/${ID}`,
                failOnStatusCode: false,
            }).then((response) =>{
                expect(response.status).to.eq(400)
                expect(response.body.status).to.eq("mongo: no documents in result")
            })
        })

        it("should return status 200 if id exist", ()=>{
            const postBody ={
                "name": "Josh Martin",
                "email": "joshmartin@gmail.com"
            }
           cy.request({
                method: "POST",
                url: "/chef/create",
                failOnStatusCode: false,
                body: JSON.stringify(postBody)
            }).then(chefCreateResponse =>{
               const id = chefCreateResponse.body["InsertedID"]
               cy.request({
                   method: "GET",
                   url: `/chef/${id}`,
                   failOnStatusCode: false,
               }).then((response) => {
                   expect(response.status).to.eq(200)
                   expect(response.body.name).to.eq(postBody.name)
               })
           })
        })
    })

    describe("/all", () =>{
        it("should return all chefs", () => {
            const postBody1 ={
                "name": "Hello Man",
                "email": "helloMan@gmail.com"
            }

            cy.request({
                method: "POST",
                url: "/chef/create",
                body: JSON.stringify(postBody1)
            }).then(()=>{
                const postBody2 ={
                    "name": "taco Man",
                    "email": "tacoMan@gmail.com"
                }
                cy.request({
                    method: "POST",
                    url: "/chef/create",
                    body: JSON.stringify(postBody2)
                }).then(()=> {
                    cy.request({
                        method: "GET",
                        url: `/chef/all`,
                    }).then((response) => {
                        const responseBody = response.body
                        expect(responseBody.some(chef => chef.name === postBody1.name)).to.eq(true)
                        expect(responseBody.some(chef => chef.name === postBody2.name)).to.eq(true)
                    })
                })
            })
        })
    })
})