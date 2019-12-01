


var spectra_files_grid;
var spectra_files_data_array = [];
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
            <img class="user-profile-icon" src="images/Reagan-Karnes.jpg"/>
        </div>

         <div class="spectra-file-attribute-layout-box">
            <div class="user-name-container">
                <span class="user-name">Reagan Karnes</span>
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
	         	<a href="javascrpt:void(0)" onclick="incremenThumbsUp()">
	             	<img class="reaction-icons" src="images/thumbs-up-512.png">
	            </a>
	             <div class="reaction-counters">12</div>
	         </div>

	         <div class="reactions-container">
	             <img class="reaction-icons" src="images/laughing-512.png">
	             <div class="reaction-counters">7</div>
	         </div>

	         <div class="reactions-container">
	             <img class="reaction-icons" src="images/happy-512.png">
	             <div class="reaction-counters">9</div>
	         </div>


	         <div class="reactions-container">
	             <img class="reaction-icons" src="images/sad-512.png">
	             <div class="reaction-counters">1</div>
	         </div>

	         <div class="reactions-container">
	             <img class="reaction-icons" src="images/angry-512.png">
	             <div class="reaction-counters">0</div>
	         </div>

	         <div class="reactions-container">
	             <img class="reaction-icons" src="images/thumbs-down-512.png">
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
	cssClass: "post-template.css",
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
		
		  spectra_files_grid = new Slick.Grid("#grid", spectra_files_data_array, spectra_files_grid_columns, slick_grid_options);
		  
		
		  spectra_files_grid.setSelectionModel(new Slick.RowSelectionModel());
	
		
		
		//Double Click Function
		//=============================================================================   
		  spectra_files_grid.onDblClick.subscribe(function (e, args) 
		  {
			  console.log("doubleClick activated");
			  console.log(e);
			  console.log(args);
	
			
  		  });
		//=============================================================================
		
		
		//Double Click Function
		//=============================================================================   
		  spectra_files_grid.onClick.subscribe(function (e, args) 
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
	
	var dataArray = new Array();

	dataArray["data_id"] = "231-xx-221";
	dataArray["filename"] = "Filename 10010bhg";
	dataArray["uploader"] = "Spectra Library";
	dataArray["upload_date"] = "2017/04/05";
	dataArray["category"] = "House Goods";
	dataArray["sample_description"] = "Windex Pro Formula";
	dataArray["solute"] = "windex";
	dataArray["solvent"] = "D.I Water";
	dataArray["concentration"] = ".002";
	dataArray["concentration_units"] = "g/cc";
	dataArray["path_length"] = "1";
	dataArray["path_length_units"] = "cm"
	dataArray["spectroscopic_method"] = "UV-VIS";
	dataArray['spectrometer_make'] = "Ocean Optics";
	dataArray['spectrometer_model'] = "HR 4000";
	dataArray['spectrometer_mode'] = "Absorbance";
	dataArray['spectrometer_resolution'] = "1";
	dataArray['spectrometer_resolution_units'] = "nm";
	dataArray['spectrometer_snr'] = "1000:1";
	dataArray['integration_time'] = "500"
	dataArray['integration_time_units'] = "ms";
	dataArray['upper_lambda'] = "1100";
	dataArray['lower_lambda'] = "190";
	dataArray['lambda_units'] = "nm";
	dataArray['additional_files'] = "background-corrected-sample-file, cuvette-background-file, solvent-background-file";
	dataArray['price'] = "109.99";
	
	
	//var dataObject = new SpectraDocumentData(dataArray);

	
	
	//var item = {file_name:dataObject, data:dataObject};

	var item = {file_name:"11", data:"11"};
	
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	/*spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	
	/*
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
	spectra_files_data_array.push(item);
*/
	
	
	spectra_files_grid.invalidate();
	spectra_files_grid.updateRowCount();
	spectra_files_grid.render();
	refreshGrid();
}

function refreshGrid()
{
	spectra_files_grid.invalidate();
	spectra_files_grid.updateRowCount();
	spectra_files_grid.render();
	console.log("Grid Refreshed!");

}
