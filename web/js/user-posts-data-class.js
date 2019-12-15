//====================================================>// JavaScript Class Start

UserPostData = (function UserPostData(constructor) //class function aka constructor
		{
			
			function UserPostData(constructor)
			{
	
				//private instance variables==================================================>
				
				//Standardized Data Variables
				//let dataID = "SS_" + Math.random().toString(36).substring(2);

				let Id = constructor['Id'];
				let content = constructor['Content'];
				let upVotes = constructor['UpVotes'];
				let downVotes = constructor['DownVotes'];
				let self = this;
				//private instance variables==================================================>
				
				//public methods====================================================>

					//file information getters
					this.getID =  function(){return Id;};
					this.getContent= function(){return content;};
					this.getUpVotes = function(){return upVotes;};
					this.getDownVotes = function(){return downVotes;};
				//public methods====================================================>
				
			}//end of cls
			
			
			return  UserPostData
			
		})();
//====================================================>// JavaScript Class End