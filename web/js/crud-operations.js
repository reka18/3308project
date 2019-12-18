function reactToPost(reaction)
{
    console.log("reacting to post");
    console.log(reaction);
    const reactionURL =  getUsername() + "/vote";
    console.log(reactionURL);
    const upId = 'upvote-' + reaction.split('-')[1];
    const downId = 'downvote-' + reaction.split('-')[1];

    console.log(upId);
    console.log(downId);

    $.ajax(
        {
            type:'GET',
            url:reactionURL,
            success: function(responseData)
            {
                if(responseData)
                {
                    const jsonData = JSON.parse(responseData);
                    document.getElementById(upId).innerHTML = jsonData['UpVotes'];
                    document.getElementById(downId).innerHTML = jsonData['DownVotes'];

                }
            },
            data: {"cast": reaction},
            cache: false,
            error: function (xhr) {
                console.log(JSON.stringify((xhr)));
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
        success: function()
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
        error: function () {
            $('.modal').click()
        }
    });
}

async function getPosts()
{
    console.log("get post fired");

    const username = getUsername();
    const postURL = username + "/post";

    let result;
    result = await $.ajax(
        {
            type: 'GET',
            url: postURL,
            success: function (responseData) {
                console.log("Post data successfully retrieved");
                console.log("Post data length: " + responseData.length);
                console.log("Post Data : " + responseData);
                return responseData.reverse()
            },
            dataType: 'json',
            data: {"limit": "500"},
            cache: false,
        });

    return result;
}

async function getUser(userName)
{
    console.log("get post fired");

    const postURL = "/user";

    let result;
    result = await $.ajax(
        {
            type: 'GET',
            url: postURL,
            success: function (responseData) {
                console.log("User data successfully retrieved");
                console.log("User data length: " + responseData.length);
                console.log("User Data : " + JSON.stringify(responseData));
                return JSON.stringify(responseData);
            },
            dataType: 'json',
            data: {"user": userName},
            cache: false,
        });

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
            success: function(responseData)
            {
                if(!responseData)
                {
                    return;
                }
                showUserSearchResults(responseData);
            },
            data:{"terms":searchTerms},
            cache: false,
            error: function (xhr)
            {
                console.log(JSON.stringify((xhr)));
            }
        });
}


function followUser(userName)
{

    const followURL = getUsername() + "/follow";

    console.log(userName);

    $.ajax(
        {
            type:'GET',
            url: followURL,
            success: function(responseData)
            {
                if(!responseData)
                {
                    return;
                }
                let cardId = '#' + userName + 'card';
                $(cardId).remove();
                refreshPosts();
            },
            data:{"user":userName},
            cache: false,
            error: function (xhr)
            {
                console.log(JSON.stringify((xhr)));
            }
        });

}

function unFollowUser(userName)
{

    const followURL = getUsername() + "/follow";

    console.log(userName);

    $.ajax(
        {
            type:'GET',
            url: followURL,
            success: function(responseData)
            {
                if(!responseData)
                {
                    return;
                }
                let cardId = '#' + userName + 'card';
                $(cardId).remove();
                refreshPosts();
            },
            data:{"user":userName, "unfollow": ""},
            cache: false,
            error: function (xhr)
            {
                console.log(JSON.stringify((xhr)));
            }
        });

}


function getUsername()
{
    let windowURL = window.location.href;
    let splitArray = windowURL.split("/");
    return splitArray[3]
}
