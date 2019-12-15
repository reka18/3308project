
var cards_holder=[];


function generate_user_card(postObject)
{
	cards_holder.push(postObject)



	const template = `
	<div class="post-layout-box-2 animated zoomInUp delay-1s">
		<div class="user-profile-icon-container-2">
			<img class="user-profile-icon-2" src="images/Reagan-Karnes-64.jpg"/>
		</div>
		<div class="spectra-file-attribute-layout-box-2">
			<div class="user-name-container-2">
				<span class="user-name-2">Reagan Karnes</span>
				<span class="user-post-date-2">2019.09.10 12:30:31pm</span>
			</div>
			<div class="user-post-container-2">
				<span class="user-post-2">${postObject.getContent()}</span>
			</div>
	
		<div class="reaction-bar-container-2">
			<div class="reaction-bar-2">
				<div class="reactions-container-2">
					<a href="javascript:void(0)" onclick="reactToPost(postID)">
						<img class="reaction-icons-2" src="images/thumbs-up-opt-512.png">
					</a>
					<div class="reaction-counters-2">12</div>
				</div>
		
				<div class="reactions-container-2">
					<a href="javascript:void(0)" onclick="reactToPost(postID,5)">
						<img class="reaction-icons-2" src="images/thumbs-down-opt-512.png">
					</a>
					<div class="reaction-counters-2">2</div>
				</div>
				
				
				<div class="reactions-container-2">
					<a href="javascript:void(0)" onclick="reactToPost(postID,5)">
						<img class="reaction-icons-2" src="images/follow-opt-512.png">
					</a>
					<div class="reaction-counters-2">2</div>
				</div>
				
				
				
			</div>
		</div>
	</div>
	`;

	return template



		/*`
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
                	<span class="user-post">${postObject.getContent()}</span>
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




		<div class="reactions-container-2">
					<a href="javascrpt:void(0)" onclick="reactToPost(postID,1)">
						<img class="reaction-icons-2" src="images/laughing-opt-512.png">
					</a>
					<div class="reaction-counters-2">7</div>
				</div>

				<div class="reactions-container-2">
					<a href="javascrpt:void(0)" onclick="reactToPost(postID,2)">
						<img class="reaction-icons-2" src="images/happy-opt-512.png">
					</a>
					<div class="reaction-counters-2">9</div>
				</div>

				<div class="reactions-container-2">
					<a href="javascrpt:void(0)" onclick="reactToPost(postID,3)">
						<img class="reaction-icons-2" src="images/sad-opt-512.png">
					</a>
					<div class="reaction-counters-2">1</div>
				</div>

				<div class="reactions-container-2">
					<a href="javascrpt:void(0)" onclick="reactToPost(postID,4)">
						<img class="reaction-icons-2" src="images/angry-opt-512.png">
					</a>
					<div class="reaction-counters-2">0</div>
				</div>

		 */


}


