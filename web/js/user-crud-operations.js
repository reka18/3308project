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



function newPost()
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


function newPost2()
{
        console.log("new post fired");
        $.ajax(
            {
                    type:'POST',
                    url: "rg/post",
                    success: function(responseData, status, responseObject)
                    {
                            //perform some action on success
                            console.log("success");

                    },
                    data: {"content" : "some text"},
                    contentType : "multipart/form-data",
                    cache: false,
                    error: function (xhr, ajaxOptions, thrownError) {
                            alert(xhr.status);
                            alert(thrownError);
                            alert(ajaxOptions);
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
