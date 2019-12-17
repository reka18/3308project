
$( window ).on("load", async function()
{

    getPosts().then(function (postData)
    {
        let cardViews = document.getElementById('grid').innerHTML;

        if(!postData)
        {
            return;
        }

        if(postData.length)
        {
            for(let x = 0; x < postData.length ; x++)
            {
                const postObject = new UserPostData(postData[x]);
                cardViews += generate_user_card(postObject);

            }
            document.getElementById('grid').innerHTML = cardViews;
        }

        else
        {}


    });


});

async function loadUserData()
{

    getThisUser().then(function (userData)
    {

        if(!userData)
        {
            return;
        }

        let data = new UserData(userData);
        document.getElementById('userNameId').innerHTML = data.getUsername();
        document.getElementById('ageId').innerHTML = data.getAge();
        document.getElementById('firstNameId').innerHTML = data.getFirstname();
        document.getElementById('lastNameId').innerHTML = data.getLastname();
        document.getElementById('emailId').innerHTML = data.getEmail();
        document.getElementById('isPublicId').innerHTML = data.getIsPublic();
        document.getElementById('friendlyDateId').innerHTML = data.getFriendlyJoinDate();
        document.getElementById('isActiveId').innerHTML = data.getIsActive();
        document.getElementById('genderId').innerHTML = data.getGender();


    });


}


function updatePosts(postsData)
{
    console.log(postsData);
    const postObject = new UserPostData(postsData[0]);
    let newCard = generate_user_card(postObject);
    $('#grid').prepend(newCard);
}

function userLogout()
{
    window.location.href = "logout";
}

function showUserSearchResults(searchResults)
{
    document.getElementById('searchCardContainer').innerHTML="";
    let jsonSearchData = JSON.parse(searchResults);
    let searchCardViews = '';


    jsonSearchData.forEach( user =>
        {
            console.log(user);
            searchCardViews += search_results_card_generator(user);
        }
    );

    $('#searchCardContainer').prepend(searchCardViews);
    $('#exampleModalLong').modal('toggle');
}



