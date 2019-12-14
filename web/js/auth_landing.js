$( window ).on("load", async function()
{
    console.log("async");

    getPosts().then(function (res)
    {
        dataView.beginUpdate();

        for(let x = 0; x < res.length ; x++)
        {

            const postObject = new UserPostData(res[x]);
            console.log(postObject.getID());
            let dataViewObject = {"Id": x, "Content": postObject}
            dataView.addItem(dataViewObject, "Id");
        }

        dataView.endUpdate();



    });


});


function updatePosts(postsData)
{

    console.log(postsData);

    dataView.beginUpdate();




    for(let x = 0; x < postsData.length  ; x++)
    {
        const postObject = new UserPostData(postsData[x]);
        console.log(postObject.getID());
        let dataViewObject = {"Id": x, "Content": postObject}
        dataView.addItem(dataViewObject, "Id");
    }

    //dataView.setItems(objectArray, "Id");
    dataView.endUpdate();

}



