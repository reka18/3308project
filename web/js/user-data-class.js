//====================================================>// JavaScript Class Start

UserData = (function UserData()
        {

            function UserData(constructor)
            {
                let id = constructor['Id'];
                let age = constructor['Age'];
                let firstname = constructor['Firstname'];
                let lastname = constructor['Lastname'];
                let email = constructor['Email'];
                let username = constructor['Username'];
                let ispublic = constructor['Public'];
                let joindate = constructor['Joindate'];
                let friendlyjoindate = constructor['FriendlyJoinDate'];
                let isactive = constructor['Active'];
                let gender = constructor['Gender'];
                let self = this;

                self.getId = function(){return id;};
                self.getAge = function(){return age;};
                self.getFirstname = function(){return firstname;};
                self.getLastname = function(){return lastname;};
                self.getEmail = function(){return email;};
                self.getUsername = function(){return username;};
                self.getIsPublic = function(){return ispublic;};
                self.getJoindate = function(){return joindate;};
                self.getFriendlyJoinDate = function(){return friendlyjoindate;};
                self.getIsActive = function(){return isactive;};
                self.getGender = function(){return gender;};
            }
            return UserData
        })();

//====================================================>// JavaScript Class End