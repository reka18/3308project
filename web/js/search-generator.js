
function search_results_card_generator(userObject)
{

	let userName = userObject['User']['Username'];
	let firstName = userObject['User']['Firstname'];
	let lastName = userObject['User']['Lastname'];
	let age = userObject['User']['Age'];
	let email = userObject['User']['Email'];
	let joinDate = userObject['User']['FriendlyJoinDate'];
	let gender = userObject['User']['Gender'];

	return `
	<div class="card text-white bg-dark mb-3 card-plus"> <!--column flex-->
		<div class="search-results-card-layout-container"> <!-- row flex -->
			<div class="search-results-image-containers">
				<img class="search-results-user-image" src="avatar?user=${userName}" alt=""/>
			</div>
			<div class="search-results-user-details-container">
				<div>
					<h5 class="card-title card-title-plus">${userName}</h5>
					<button style="margin-left: 350px;">
						<img src="images/follow-opt-w-512.png" id="${userName}card" class="search-results-follow-icon followMe rotate-" onclick="followUser('${userName}');" alt=""/>
					</button>
			
				</div>
				<div class="search-results-user-information-container">
					<span class="search-results-user-information">Firstname: ${firstName}</span>
					<span class="search-results-user-information">Lastname: ${lastName} </span>
					<span class="search-results-user-information">Age: ${age}</span>
					<span class="search-results-user-information">Email: ${email}</span>
					<span class="search-results-user-information">Join Date: ${joinDate}</span>
					<span class="search-results-user-information">Gender: ${gender}</span>
				</div>
			</div>
		</div>
	</div>
	`;

}


