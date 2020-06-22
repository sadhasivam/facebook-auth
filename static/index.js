window.fbAsyncInit = function () {
  FB.init({
    appId: "715153329275843",
    cookie: true,
    xfbml: true,
    version: "v7.0",
  });

  FB.AppEvents.logPageView();
};

(function (d, s, id) {
  var js,
    fjs = d.getElementsByTagName(s)[0];
  if (d.getElementById(id)) {
    return;
  }
  js = d.createElement(s);
  js.id = id;
  js.src = "https://connect.facebook.net/en_US/sdk.js";
  fjs.parentNode.insertBefore(js, fjs);
})(document, "script", "facebook-jssdk");

const statusChangeCallback = (response) => {
  logGutter(`<h3> FaceBook Login Status </h3> <pre> ${JSON.stringify(response)}</pre>`);
  enableDisableButton((response.status === "connected"));
  
};

const signOut = () => {
  FB.logout(function (response) {
    logGutter(`<h3> FaceBook Logout Status </h3> <pre> ${JSON.stringify(response)}</pre>`);
    enableDisableButton(false)
    // Person is now logged out
  });
};

const enableDisableButton = (loggedIn) => {
  if (loggedIn) {
    document.getElementById("fbLogin").style.display = "none";
    document.getElementById("signOut").style.display = "block";
  } else {
    document.getElementById("fbLogin").style.display = "block";
    document.getElementById("signOut").style.display = "none";
  }
};

const logGutter = (msg) => {
  const elem = document.getElementById("logGutter");
  elem.innerHTML = elem.innerHTML + msg;
};
function checkLoginState() {
  FB.getLoginStatus(function (response) {
    statusChangeCallback(response);
  });
}

window.onload = () => {
  FB.getLoginStatus(function (response) {
    statusChangeCallback(response);
  });
};
