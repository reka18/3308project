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
				let uploadDate = constructor['upload_date'];
				let thumbsUpResponse = constructor['thumbs_up_response'];
				let laughingResponse = constructor['laughing_response'];
				let happyResponse = constructor['happy_response'];
				let sadResponse = constructor['sad_response'];
				let angryResponse = constructor['angry_response'];
				let thumbsDownResponse = constructor['thumbs_down_response'];
				let self = this;
				//private instance variables==================================================>
				
				//public methods==================================================>
				
				
					//file information getters
					this.getPostID = function(){return postID;};
					this.getUsername = function(){return username;};
					this.getPostDate = function(){return PostDate;};
					this.getThumbsUpResponse = function(){return uploadDate;};
				//public methods==================================================>
				
				
				
				
	
				
			}//end of cls
			
			
			return  UserPostData
			
			
		
		})();
//====================================================>// JavaScript Document