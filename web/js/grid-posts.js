


var user_posts_grid;
var user_posts_data_array = [];
var pageNumber_5 = 1;
var pageLimit_5 = null;
var searchKey_5 = '';




/*
function formatter_5(row, cell, value, columnDef, dataContext) 
		 {
			 console.log(row);
			 console.log(cell);
			 console.log(value);
			 console.log(columnDef);
			 console.log(dataContext);
			 console.log(grid_5.getCellNode(row,cell));
			 
			 
			 var dataHolder = dataContext.Data;
			 
			 
			 
			 if(dataHolder.getErrors())
			 {
				  return '<div id="icon_rs_' + row + '" class="cell-icon-holder cell-icon-error"></div>';
				 
			 }
			 
			 else
			 {
			 
			 
				 if(dataHolder.getViewed() == 0)
				 {
					 return '<div id="icon_rs_' + row + '" class="cell-icon-holder cell-icon-unviewed"></div>';
					  
				 }
				 
				 else
				 {
					  return '<div id="icon_rs_' + row + '" class="cell-icon-holder cell-icon-results-file-viewed"></div>';
					 
				 }
			 }
			 
			
			 
        	return value;
		}	  
		
		
		*/
		
		
		function requiredFieldValidator(value) {
		  if (value == null || value == undefined || !value.length) {
			  console.log("cant be empty neeegrooooo!");
			return {valid: false, msg: "This is a required field"};
		  } else {
			return {valid: true, msg: null};
		  }
		}
		

/*
function detailButton(row, cell, value, columnDef, dataContext)
{
	 
	 return '<button class="slick-cell-button">Details</button>'
	 	
}
*/


/*
function resultsFile(row, cell, value, columnDef, dataContext)
{
	
	var filename = value.getFilename();
	var method = value.getSpectroscopicMethod();
	var mode = value.getSpectrometerMode();
	var range = value.getLowerLambda() + ' - ' + value.getUpperLambda() + value.getLambdaUnits();
	var resolution = value.getSpectrometerResolution() + value.getSpectrometerResolutionUnits();
	var snr = value.getSpectrometerSNR();
	
	return '\
	\
	\
	<div class="spectra-file-layout-box">\
	\
	\
	<div class="spectra-file-icon-layout-box"><img class="spectra-file-icon" src="images/spectra-library-file-icon2.png"/></div>\
    \
    \
    <div class="spectra-file-attribute-layout-box">\
        <div class="spectra-file-attribute"><span class="spectra-file-attribute-header">Filename: </span><span class="spectra-file-attribute-description">'+filename+'</span></div>\
        <div class="spectra-file-attribute"><span class="spectra-file-attribute-header">Spectroscopic Method: </span><span class="spectra-file-attribute-description">'+method+'</span></div>\
		<div class="spectra-file-attribute"><span class="spectra-file-attribute-header">Mode: </span><span class="spectra-file-attribute-description">'+mode+'</span></div>\
        <div class="spectra-file-attribute"><span class="spectra-file-attribute-header">Wavelength Range: </span><span class="spectra-file-attribute-description">'+range+'</span></div>\
        <div class="spectra-file-attribute"><span class="spectra-file-attribute-header">Resolution: </span><span class="spectra-file-attribute-description">'+resolution+'</span></div>\
       	<div class="spectra-file-attribute"><span class="spectra-file-attribute-header">Signal to Noise: </span><span class="spectra-file-attribute-description">'+snr+'</span></div>\
       	<div class="spectra-file-attribute"><span class="spectra-file-attribute-header">Signal to Noise: </span><span class="spectra-file-attribute-description">&nbsp;</span></div>\
        <div class="spectra-file-attribute"><button class="file-details-button" onclick="someFunction('+row+')"><span class="spectra-file-attribute-description">Load Details</span></button></div>\
    </div>\
	\
	\
	</div>';
	
}
*/



function resultsFile(row, cell, value, columnDef, dataContext)
{




	return `
	<div class="post-layout-box">
        <div class="user-profile-icon-container">
            <img class="user-profile-icon" src="images/Reagan-Karnes-64.jpg"/>
        </div>

         <div class="spectra-file-attribute-layout-box">
            <div class="user-name-container">
                <span class="user-name">Reagan Karnes</span>
                <span class="user-post-date">2019.09.10 12:30:31pm</span>
            </div>
            <div class="user-post-container">
                <span class="user-post">
                </span>
            </div>
        </div>
        </div>
        <div class="reaction-bar-container">
         <div class="reaction-bar">
	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID)">
	             	<img class="reaction-icons" src="images/thumbs-up-opt-512.png">
	            </a>
	             <div class="reaction-counters">12</div>
	         </div>

	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID,1)">
	         		<img class="reaction-icons" src="images/laughing-opt-512.png">
	         	</a>
	             <div class="reaction-counters">7</div>
	         </div>

	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID,2)">
	            	<img class="reaction-icons" src="images/happy-opt-512.png">
	            </a>
	             <div class="reaction-counters">9</div>
	         </div>


	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID,3)">
	         		<img class="reaction-icons" src="images/sad-opt-512.png">
	            </a>
	             <div class="reaction-counters">1</div>
	         </div>

	         <div class="reactions-container">
	         	<a href="javascrpt:void(0)" onclick="reactToPost(postID,4)">
	            	<img class="reaction-icons" src="images/angry-opt-512.png">
	            </a>
	             <div class="reaction-counters">0</div>
	         </div>

	         <div class="reactions-container">
	         <a href="javascrpt:void(0)" onclick="reactToPost(postID,5)">
	             <img class="reaction-icons" src="images/thumbs-down-opt-512.png">
	         </a>
	             <div class="reaction-counters">2</div>
	         </div>
         </div>
     </div>
    </div>`;
	
}





var spectra_files_grid_columns = 
[
  
  {
	id: "file_name",
	name: "Results",
	field: "file_name",
	width: 10,
	minWidth: 0,
	cssClass: "post-template-styling.css",
	headerCssClass:"spectra-files-slick-grid-header",
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
		
		  user_posts_grid = new Slick.Grid("#grid", user_posts_data_array, spectra_files_grid_columns, slick_grid_options);
		  
		
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
		
		
		
		  addItem();
		
		});
		
		
	//END Grid Setup
	//=============================================================================


//column names		  
//file-name		  
//spectoscopic-method
//mode
//wavelength-range
//resolution
//details
//data

function addItem()
{

	var dataObject = new UserPostData(ajaxResponseDataSimulation());

	var grid_item = {file_name:dataObject, data:dataObject};
	
	user_posts_data_array.push(grid_item);
	user_posts_data_array.push(grid_item);
	user_posts_data_array.push(grid_item);
	user_posts_data_array.push(grid_item);
	user_posts_data_array.push(grid_item);
	user_posts_data_array.push(grid_item);
	user_posts_data_array.push(grid_item);
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
	user_posts_grid.invalidate();
	user_posts_grid.updateRowCount();
	user_posts_grid.render();
	console.log("Grid Refreshed!");

}
