    console.log('start')
    var currentCount = 5
    
    while (currentCount > 0)
    {
        // create request:
        const postRequest = {
            url: 'localhost:3000/user',
            method: 'POST',
            header: {
                'Content-Type': 'application/json',
                //'Authorization': 'Bearer no'
            },
            body: {
                mode: 'raw',
                raw: JSON.stringify({
                    "Name": "Bob " + currentCount,
                    "Email": "test@email" + currentCount, })
            }
        };

        pm.sendRequest(postRequest, function (err, response) {
            console.log(response);
            console.log('run' + currentCount + ' ' + response.responseTime);
            //setTimeout(() => {}, 15000); 
        });
        currentCount--;
    }
    console.log("end")