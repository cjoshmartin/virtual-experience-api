context("Experiences", () => {

   describe("/create", () =>{
      it('should return a 400 if not given a chefID', () => {
         const postBody = {
            "name": "Dumpling Making",
            "description": "Come learn with our master chefs and learn how to make dumplings",
            "price": 100.0,
         }

         cy.request({
            method: "POST",
            url: "/experience/create",
            failOnStatusCode: false,
            body: JSON.stringify(postBody)
         }).then((response) =>{
            expect(response.status).to.eq(400)
         })
      })

      it('should return status code 200 if body is valid', () => {
         let chefID;
         const chefBody ={
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
            const postBody= {
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
            }).then((response) =>{
               expect(response.status).to.eq(200)
            })
         })
      })

      it('should return status code 400 if price is missing', () => {
         let chefID;
         const chefBody ={
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
            const postBody= {
               "name": "Dumpling Making",
               "description": "Come learn with our master chefs and learn how to make dumplings",
               "chefid": chefID
            }
            cy.request({
               method: "POST",
               url: "/experience/create",
               failOnStatusCode: false,
               body: JSON.stringify(postBody)
            }).then((response) =>{
               expect(response.status).to.eq(400)
            })
         })
      })
   })
    describe("/:id", ()=>{
       it("should return status 400 if id is invalid", ()=>{
            const id = "60d7746e4ee6efaaae06cac8"

          cy.request({
             method: "GET",
             url: `/user/${id}`,
             failOnStatusCode: false,
          }).then((response) =>{
             expect(response.status).to.eq(400)
          })
       })
       it("should find record by id if id is valid", () => {

          let chefID;
          const chefBody ={
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
             const postBody= {
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
             }).then((response) =>{
                cy.request({
                   method: "GET",
                   url: `/experience/${response.body["InsertedID"]}`
                }).then(rep => {
                   expect(rep.status).to.eq(200)
                })
             })
          })
       })
    })
})
