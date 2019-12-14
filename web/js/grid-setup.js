


var user_posts_grid;
var user_posts_data_array = [];
dataView = new Slick.Data.DataView({inlineFilters: true});
var pageNumber_5 = 1;
var pageLimit_5 = null;
var searchKey_5 = '';



function requiredFieldValidator(value) {
	if (value == null || value == undefined || !value.length) {
		console.log("cant be empty neeegrooooo!");
		return {valid: false, msg: "This is a required field"};
	} else {
		return {valid: true, msg: null};
	}
}

function resultsFile(row, cell, value, columnDef, dataContext)
{
	console.log("results file value object is..." + dataContext);

	const template = `
		<div class="post-layout-box">
			<div class="user-profile-icon-container">
            	<img class="user-profile-icon" src="../images/Reagan-Karnes-64.jpg"/>
        	</div>
        	<div class="spectra-file-attribute-layout-box">
            	<div class="user-name-container">
                	<span class="user-name">Reagan Karnes</span>
                	<span class="user-post-date">2019.09.10 12:30:31pm</span>
                </div>
                <div class="user-post-container">
                	<span class="user-post">${value.getContent()}</span>
                </div>
            </div>
         </div>
       <div class="reaction-bar-container">
         <div class="reaction-bar">
	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID)">
	             	<img class="reaction-icons" src="../images/thumbs-up-opt-512.png">
	            </a>
	             <div class="reaction-counters">12</div>
	         </div>

	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID,1)">
	         		<img class="reaction-icons" src="../images/laughing-opt-512.png">
	         	</a>
	             <div class="reaction-counters">7</div>
	         </div>

	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID,2)">
	            	<img class="reaction-icons" src="../images/happy-opt-512.png">
	            </a>
	             <div class="reaction-counters">9</div>
	         </div>
	         
	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID,3)">
	         		<img class="reaction-icons" src="../images/sad-opt-512.png">
	            </a>
	             <div class="reaction-counters">1</div>
	         </div>

	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID,4)">
	            	<img class="reaction-icons" src="../images/angry-opt-512.png">
	            </a>
	             <div class="reaction-counters">0</div>
	         </div>

	         <div class="reactions-container">
	         <a href="javascrpt:void(0)" onclick="reactToPost(postID,5)">
	             <img class="reaction-icons" src="../images/thumbs-down-opt-512.png">
	         </a>
	             <div class="reaction-counters">2</div>
	         </div>
         </div>
     </div>
	`;



	return template







}





var grid_columns_setup =
	[

		{
			id: "Content",
			name: "Results",
			field: "Content",
			width: 10,
			minWidth: 0,
			editor: Slick.Editors.Text,
			validator: requiredFieldValidator,
			formatter: resultsFile,
		},



		{
			id: "data",
			field: "Data",
			name:"data",
			width: 0,
			height:0,
			maxWidth: 0,
			minWidth: 0,
			cssClass: "hidden-cells",
			headerCssClass: "hidden-header",

		}

	];


var slick_grid_options = {
	editable: false,
	enableAddRow: false,
	enableCellNavigation: true,
	//rowHeight: 220, /* was 187*/
	rowHeight: 260, /* was 187*/
	syncEditorLoading:true,
	//requires double click to edit??
	autoEdit: false,

	enableColumnReorder: false,
	autoHeight: false,  //This disables vertical scrolling.
	forceFitColumns:true,
	fullWidthRows:true,
	rerenderOnResize:true,
};



//Grid Setup
//=============================================================================
$(function ()
{
	dataView.setItems([], "Id");

	user_posts_grid = new Slick.Grid("#grid", dataView, grid_columns_setup, slick_grid_options);


	user_posts_grid.setSelectionModel(new Slick.RowSelectionModel());



	//Double Click Function
	//=============================================================================
	user_posts_grid.onDblClick.subscribe(function (e, args)
	{
		console.log("doubleClick activated");
		console.log(e);
		console.log(args);


	});
	//=============================================================================


	//Double Click Function
	//=============================================================================
	user_posts_grid.onClick.subscribe(function (e, args)
	{
		console.log("Click activated");
		console.log(e);
		console.log(args);



	});
	//=============================================================================

	// Make the grid respond to DataView change events.
	dataView.onRowCountChanged.subscribe(function (e, args) {
		console.log('count');
		console.log(e);
		console.log(args);
		user_posts_grid.updateRowCount();
		user_posts_grid.render();
		dataView.refresh();
		refreshGrid();
	});

	dataView.onRowsChanged.subscribe(function (e, args) {
		console.log('rows');
		console.log(e);
		console.log(args);
		user_posts_grid.invalidateRows(args.rows);
		user_posts_grid.render();
		dataView.refresh();
		refreshGrid();
	});

});


//END Grid Setup
//=============================================================================


function addItem()
{

	let dataObject = new UserPostData(ajaxResponseDataSimulation());

	let grid_item = {file_name:dataObject, data:dataObject};

	for(let x = 0; x < 10 ; x++)
	{
		user_posts_data_array.push(grid_item);

	}


	refreshGrid();
}

function ajaxResponseDataSimulation()
{
	let dataArray = new Array();

	dataArray["post_id"] = "231-xx-221";
	dataArray["username"] = "Regan Karnes";
	dataArray["profile_picture"] = "images/Reagan-Karnes.jpg";
	dataArray["post_date"] = "2019/04/05";
	dataArray["thumbs_up"] = "5";
	dataArray["laughing"] = "12";
	dataArray["happy"] = "9";
	dataArray["sad"] = "11";
	dataArray["angry"] = "7";
	dataArray["thumbs_down"] = "7";

	return dataArray;

}

function refreshGrid()
{
	user_posts_grid.invalidateAllRows();
	user_posts_grid.updateRowCount();
	user_posts_grid.render();
	console.log("Grid Refreshed!");

}

function clearGridData()
{
	user_posts_data_array = [];

}
