
function search_results_card_generator(userObject)
{
	var userName = userObject['User']['Username'];

	return `
	 <div class="card text-white bg-dark mb-3 card-plus"> <!--column flex-->
            <div class="search-results-card-layout-container"> <!-- row flex -->
                <div class="search-results-image-containers">
                	<img class="search-results-user-image" src="avatar?user=${userName}"/>
                </div>
                <div class="search-results-user-details-container">
                    <div><h5 class="card-title card-title-plus">${userName}</h5><a href="#"><img class="search-results-follow-icon followMe" onclick="followUser('userObject[User][Username]')" class="search-results-follow-icon" src="images/follow-opt-w-512.png"/></a></div>
                    <div class="search-results-user-information-container">
                        <span class="search-results-user-information">First name: ${userObject['User']['Firstname']} </span>
                        <span class="search-results-user-information">Last name: ${userObject['User']['Lastname']} </span>
                        <span class="search-results-user-information">Age: ${userObject['User']['Age']}</span>
                        <span class="search-results-user-information">Email:  ${userObject['User']['Email']}</span>
                        <span class="search-results-user-information">Join Date: ${userObject['User']['Joindate']}</span>
                        <span class="search-results-user-information">Gender: ${userObject['User']['Gender']}</span>
                    </div>
                </div>
            </div>
        </div>
	`;

	/*
	<div class="reactions-container-2" onclick=reactToPost("up-"+${postObject.getId()})>
		<img class="reaction-icons-2" src="images/thumbs-up-opt-512.png">
		<div class="reaction-counters-2">${postObject.getUpVotes()}</div>
	</div>
	 */

}


