
console.log("index.js loaded");

$( "#login-form" ).submit(function(event)
{
  event.preventDefault();

  console.log("redirecting...");
   
  //window.location.replace("https://peanuts.com");
  window.location.replace('/views/user_account.html');
});


/*
var example2 = new Vue({
  el: '#login-controllers',
  data: {
    name: 'Vue.js'
  },
  // define methods under the `methods` object
  methods: {
    greet: function (event) {
      // `this` inside methods points to the Vue instance
      alert('Welcome back! Make sure you\'ve filled out all login fields');
      // `event` is the native DOM event
    }
  }
})
*/
