let map
let allConcerts
let geocoder

// Initialize and add the map
function initMap() {
  allConcerts = document.getElementsByClassName("concert-place");
  geocoder = new google.maps.Geocoder();

  // The map, centered at exact place
  map = new google.maps.Map(document.getElementById("map"), {
    center: new google.maps.LatLng(0, 0),
    zoom: 1,
  });

  LoadMarks()
}

function LoadMarks() {
  let x = 0

  var markerInterval = setInterval(function () {
    console.log(x)
    if (x < allConcerts.length) {
      console.log(allConcerts.length)
      var address = allConcerts[x].innerHTML
      geocoder.geocode({ 'address': address }, function (results, status) {
        if (status == google.maps.GeocoderStatus.OK) {
          const concertPlace = { lat: results[0].geometry.location.lat(), lng: results[0].geometry.location.lng() };

          new google.maps.Marker({
            position: concertPlace,
            map: map,
          });
        } else {
          console.log(status)
        }
      });
    } else {
      clearInterval(markerInterval);
    }

    x++;
  }, 450);
}

function GoToConcert(address) {
  var geocoder = new google.maps.Geocoder();
  geocoder.geocode({ 'address': address }, function (results, status) {
    if (status == google.maps.GeocoderStatus.OK) {
      const concertPlace = { lat: results[0].geometry.location.lat(), lng: results[0].geometry.location.lng() };

      map.setCenter(concertPlace);
      map.setZoom(7);
    }
  });
}