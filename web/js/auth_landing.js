$( window ).on("load", async function()
{

    getPosts().then(function (postData)
    {
        let cardViews = document.getElementById('grid').innerHTML;

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


function updatePosts(postsData)
{
    console.log(postsData);
    const postObject = new UserPostData(postsData[0]);
    let newCard = generate_user_card(postObject)
    $('#grid').prepend(newCard);
}



