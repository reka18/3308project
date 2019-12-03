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
            cache: false,
            error: /* some errorfunction */

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
            cache: false,
            error: /* some errorfunction */

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
            cache: false,
            error: /* some errorfunction */

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
            cache: false,
            error: /* some errorfunction */

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
            error: /* some errorfunction */

        });
}
