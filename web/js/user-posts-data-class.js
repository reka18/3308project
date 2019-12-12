//====================================================>// JavaScript Class Start

UserPostData = (function UserPostData(constructor) //class function aka constructor
		{
			
			function UserPostData(constructor)
			{
	
				//private instance variables==================================================>
				
				//Standardized Data Variables
				//let dataID = "SS_" + Math.random().toString(36).substring(2);
				let postID = constructor['post_id'];
				let username = constructor['username'];
				let profilePicture = constructor['profile_picture']
				let postDate = constructor['post_date'];
				let thumbsUpResponse = constructor['thumbs_up'];
				let laughingResponse = constructor['laughing'];
				let happyResponse = constructor['happy'];
				let sadResponse = constructor['sad'];
				let angryResponse = constructor['angry'];
				let thumbsDownResponse = constructor['thumbs_down'];
				let self = this;
				//private instance variables==================================================>
				
				//public methods====================================================>

					//file information getters
					this.getPostID = function(){return postID;};
					this.getUsername = function(){return username;};
					this.getProfilePicture = function(){return profilePicture;};
					this.getPostDate = function(){return postDate;};
					this.getThumbsUpResponse = function(){return thumbsUpResponse;};
					this.getLaughingResponse = function(){return laughingResponse;};
					this.getHappyResponse = function(){return happyResponse;};
					this.getSadResponse = function(){return sadResponse;};
					this.getAngryResponse = function(){return angryResponse;};
					this.getThumbsDownResponse = function(){return thumbsDownResponse;};
				//public methods====================================================>
				
			}//end of cls
			
			
			return  UserPostData
			
		})();
//====================================================>// JavaScript Class End