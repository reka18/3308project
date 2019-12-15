//====================================================>// JavaScript Class Start

UserPostData = (function UserPostData(constructor) //class function aka constructor
		{
			
			function UserPostData(constructor)
			{
	
				//private instance variables==================================================>
				
				//Standardized Data Variables
				//let dataID = "SS_" + Math.random().toString(36).substring(2);

				let username = constructor['UserName'];
				let content = constructor['Content'];
				let upvotes = constructor['UpVotes'];
				let downvotes = constructor['DownVotes'];
				let date = constructor['Date'];
				let self = this;
				//private instance variables==================================================>
				
				//public methods====================================================>
					//file information getters
					self.getUserName = function(){return username;};
					self.getContent= function(){return content;};
					self.getUpVotes = function(){return upvotes;};
					self.getDownVotes = function(){return downvotes;};
					self.getDate = function(){return date;};
				//public methods====================================================>
				
			}//end of cls
			
			
			return  UserPostData
			
		})();
//====================================================>// JavaScript Class End