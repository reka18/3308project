
var cards_holder=[];


function generate_user_card(postObject)
{
	cards_holder.push(postObject);
	const username = getUsername();

	return `
	<div class="post-layout-box-2 animated zoomInUp delay-1s">
		<div class="user-profile-icon-container-2">
			<img class="user-profile-icon-2 w-75 h-25" src="avatar?user=${postObject.getUserName()}" style="margin:10%;"/>
		</div>
		<div class="spectra-file-attribute-layout-box-2">
			<div class="user-name-container-2" style="margin-top:5px">
				<span class="user-name-2">${postObject.getUserName()}</span>
				<span class="user-post-date-2">${postObject.getFriendlyDate()}</span>
			</div>
			<div class="user-post-container-2">
				<span class="user-post-2">${postObject.getContent()}</span>
			</div>
	
		<div class="reaction-bar-container-2">
			<div class="reaction-bar-2">
				<div class="reactions-container-2">
					<a href="/${username}/vote?cast=up-${postObject.getId()}">
						<img class="reaction-icons-2" src="images/thumbs-up-opt-512.png">
					</a>
					<div class="reaction-counters-2">${postObject.getUpVotes()}</div>
				</div>
		
				<div class="reactions-container-2">
					<a href="/${username}/vote?cast=down-${postObject.getId()}">
						<img class="reaction-icons-2" src="images/thumbs-down-opt-512.png">
					</a>
					<div class="reaction-counters-2">${postObject.getDownVotes()}</div>
				</div>
			</div>
		</div>
	</div>
	`;
}


