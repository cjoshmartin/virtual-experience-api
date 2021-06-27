

context("Order Endpoints", () => {
    let chefID, experienceID, orderID;
    it('should return 400 error if no experienceID or no ChefID is provided', ()=> {
        const postBody = {
            "tip": 10.0,
            "taxes": 0.10,
            "subTotal": 10.00
        }
        cy.request({
            method: "POST",
            url: "/order/create",
            failOnStatusCode: false,
            body: JSON.stringify(postBody)
        } )
            .then((response) => {
                expect(response.status).to.eq(400)
            })
    })
    it('should return status 200 if  body is valid', ()=> {
        const chefBody = {
            "name": "Josh Martin",
            "email": "joshmartin@gmail.com"
        }
        cy.request({
            method: "POST",
            url: "/chef/create",
            failOnStatusCode: false,
            body: JSON.stringify(chefBody)
        }).then(chefCreateResponse => {
            chefID = chefCreateResponse.body["InsertedID"]
            const postBody = {
                "name": "Dumpling Making",
                "description": "Come learn with our master chefs and learn how to make dumplings",
                "price": 100.0,
                "chefid": chefID
            }

            cy.request({
                method: "POST",
                url: "/experience/create",
                failOnStatusCode: false,
                body: JSON.stringify(postBody)
            }).then((experienceResponse) => {
                experienceID = experienceResponse.body["InsertedID"]
                const postBody = {
                    "chefid": chefID,
                    "experienceid": experienceID,
                    "tip": 10.0,
                    "taxes": 0.10,
                    "subTotal": 10.00
                }
                cy.request({
                    method: "POST",
                    url: "/order/create",
                    failOnStatusCode: false,
                    body: JSON.stringify(postBody)
                }).then((response) => {
                    expect(response.status).to.eq(200)
                })
            })
        })
    })
    })