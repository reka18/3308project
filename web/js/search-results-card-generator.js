
var search_results_holder=[];


function search_results_card_generator(searchResultsObject)
{
	search_results_holder.push(postObject);

	return `
	 <div class="card text-white bg-dark mb-3 card-plus"> <!--column flex-->
            <div class="search-results-card-layout-container"> <!-- row flex -->
                <div class="search-results-image-containers">
                	<img class="search-results-user-image" src="avatar?user=rg"/>
                </div>
                <div class="search-results-user-details-container">
                    <h5 class="card-title card-title-plus">RK Killer</h5>
                    <div class="search-results-user-information-container">
                        <span">First name: Regan</span>
                        <span style="display:block; ">Last name: Karnes</span>
                        <span style="display:block; ">Age: 100</span>
                        <span style="display:block; word-break: break-word; ">Email: Regan.Karnes@plutoemailsystem.net.woa.net.org</span>
                        <span style="display:block; ">Join Date: 2019-12-19</span>
                        <span style="display:block; ">Gender: M</span>
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


