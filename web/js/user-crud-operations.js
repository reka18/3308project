function reactToPost()
{
    $.ajax(
        {
            type:'POST',
            url: "some/endpoint",
            success: function(responseData, status, responseObject)
            {
                //perform some action on success

            },
            dataType: 'json',
            data: {/*data that we will be passing to the end point*/},
            cache: false

        });
}


function deletePost()
{
    $.ajax(
        {
            type:'DELETE',
            url: "some/endpoint",
            success: function(responseData, status, responseObject)
            {
                //perform some action on success

            },
            dataType: 'json',
            data: {/*data that we will be passing to the end point*/},
            cache: false

        });
}


function newPost2() //TODO Why does this not let you change it to newPost? Why the 2?
{
        console.log("new post fired");

        let windowURL = window.location.href;
        let splitArray = windowURL.split("/");
        const username = splitArray[3];
        const postURL = username + "/post";



        const postData = document.getElementById('postText').value;

        if(postData === "")
        {
            alert("Enter some text to post!");
            return;
        }


        $.ajax(
            {
                    type:'POST',
                //TODO grab un from url onload and save for crud operations
                    url: postURL,
                    success: function(responseData, status, responseObject)
                    {
                            //perform some action on success
                            console.log("success");
                            console.log(responseData);
                            console.log(status);
                            console.log(responseObject);
                            console.log("End Transmission5");
                            $('.modal').click()

                    },
                    data: {

                        "Content-Type": "text/html; charset=utf-8",
                        "content": postData},
                    dataType: 'html',
                    cache: false,
                    error: function (xhr, ajaxOptions, thrownError) {
                            alert(xhr.status);
                            alert(thrownError);
                            alert(ajaxOptions);
                            $('.modal').click()

                    }
            });
}



function replyToPost()
{
    $.ajax(
        {
            type:'POST',
            url: "some/endpoint",
            success: function(responseData, status, responseObject)
            {
                //perform some action on success

            },
            dataType: 'json',
            data: {/*data that we will be passing to the end point*/},
            cache: false

        });
}




function editPost()
{
    $.ajax(
        {
            type:'PUT',
            url: "some/endpoint",
            success: function(responseData, status, responseObject)
            {
                //perform some action on success

            },
            dataType: 'json',
            data: {/*data that we will be passing to the end point*/},
            cache: false,


        });
}
