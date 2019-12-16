function reactToPost(reaction)
{
    console.log("reacting to post");
    console.log(reaction);
    const reactionURL =  getUsername() + "/vote";
    console.log(reactionURL);

    $.ajax(
        {
            type:'GET',
            url:reactionURL,
            success: function(responseData, status, responseObject)
            {
                console.log("Success");
                console.log(responseData);
                console.log(status);
                console.log(responseObject);
            },
            data: {"cast": reaction},
            cache: false,
            error: function (xhr, ajaxOptions, thrownError) {
                console.log(JSON.stringify((xhr)));
                alert(xhr);
                alert(thrownError);
                alert(ajaxOptions);
            }
        });
}

function newPost()
{
        const postURL = getUsername()+ "/post";
        const postData = document.getElementById('postText').value;
        document.getElementById("postText").innerText="";

        if(postData === "")
        {
            alert("Enter some text to post!");
            return;
        }

        $.ajax({
                    type:'POST',
                    url: postURL,
                    success: function(responseData, status, responseObject)
                    {
                            $('.modal').click();
                            getPosts().then(function (results)
                            {
                                updatePosts(results);
                            });

                    },
                    data: {"Content-Type": "text/html; charset=utf-8", "content": postData},
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

async function getPosts()
{
    console.log("get post fired");

    const username = getUsername();
    const postURL = username + "/post";

    const [result] = await Promise.all([$.ajax(
        {
            type: 'GET',
            url: postURL,
            success: function (responseData, status, responseObject) {
                console.log("Post data successfully retrieved");
                console.log("Post data length: " + responseData.length);
                console.log("Post Data : " + responseData);
                return responseData.reverse()
            },
            dataType: 'json',
            data: {"limit": "500"},
            cache: false,
            error: function (xhr, ajaxOptions, thrownError) {
                alert(xhr.status);
                if (!thrownError) {
                    alert(xhr.status);
                    alert(thrownError);
                    alert(ajaxOptions);
                }
            }
        })]);

    return result;
}




function userSearch()
{
    let getUrl = window.location;
    let baseUrl = getUrl .protocol + "//" + getUrl.host + "/";
    const searchURL = baseUrl + "search";
    let searchTerms = document.getElementById('userSearchBar').value;

    if(!searchTerms) {
        return}

    $.ajax(
        {
            type:'GET',
            url: searchURL,
            success: function(responseData, status, responseObject)
            {
                console.log("Search: " + JSON.stringify(responseData));
            },
            data:{"terms":searchTerms},
            dataType: 'json',
            cache: false,
            error: function (xhr, ajaxOptions, thrownError) {
                console.log(JSON.stringify((xhr)));
                alert(xhr);
                alert(thrownError);
                alert(ajaxOptions);
            }


        });
}


function getUsername()
{
    let windowURL = window.location.href;
    let splitArray = windowURL.split("/");
    return splitArray[3]

}

function getBaseUrl()
{
    let getUrl = window.location;
    let baseUrl = getUrl .protocol + "//" + getUrl.host + "/";
    return baseUrl;
}
