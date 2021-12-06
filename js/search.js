function CheckStyle() {
  var groupSection = document.getElementById("groups-section");
  var memberSection = document.getElementById("members-section");
  var locationSection = document.getElementById("locations-section");

  sectionArray = [groupSection, memberSection, locationSection]

  for (var i = 0; i < 3; i++) {
    var toDisplay = true
    for (var k = 0; k < sectionArray[i].getElementsByTagName("li").length; k++) {
      if (sectionArray[i].getElementsByTagName("li")[k].style.display == "none") {
        toDisplay = false
        continue
      } else {
        toDisplay = true
        break
      }
    }

    if (toDisplay) {
      sectionArray[i].style.display = "block";
    } else {
      sectionArray[i].style.display = "none";
    }
  }
}

function ShowData() {
  var input, filter, ul, li, a, i, txtValue, selection, exactSearch;
  selection = document.getElementById('category-selector').value;
  input = document.getElementById("myInput");
  filter = input.value.toUpperCase();
  ul = document.getElementById("myUL");
  li = ul.getElementsByTagName("li");

  (selection != "Artist/Band") ? document.getElementById("exact-search-checkbox").disabled = true : document.getElementById("exact-search-checkbox").disabled = false

  exactSearch = document.getElementById("exact-search-checkbox").checked;
  for (i = 0; i < li.length; i++) {
    a = li[i].getElementsByTagName("a")[0];
    if (a.attributes.length != 2) {
      var attributesValues = [a.attributes[2].value, a.attributes[3].value, a.attributes[4].value, a.attributes[5].value]
    }

    txtValue = a.textContent || a.innerText;

    if (selection == "Artist/Band") {
      if (a.attributes[0].value == "groups") {
        if (exactSearch) {
          txtValue.toUpperCase().indexOf(filter) == 0 ? li[i].style.display = "" : li[i].style.display = "none";
          continue
        }

        txtValue.toUpperCase().indexOf(filter) > -1 ? li[i].style.display = "" : li[i].style.display = "none";
        continue
      }
      else {
        li[i].style.display = "none";
        continue
      }
    }
    else if (selection == "Members") {
      if (a.attributes[0].value == "groups" || a.attributes[0].value == "members") {
        if (attributesValues != null) {
          attributesValues[3].toString().toUpperCase().indexOf(filter) > -1 ? li[i].style.display = "" : li[i].style.display = "none"
          continue
        }

        txtValue.toUpperCase().indexOf(filter) > -1 ? li[i].style.display = "" : li[i].style.display = "none";
        continue
      }
      else {
        li[i].style.display = "none";
        continue
      }
    }
    else if (selection == "Locations") {
      if (a.attributes[0].value == "groups") {
        attributesValues[2].toString().toUpperCase().indexOf(filter) > -1 ? li[i].style.display = "" : li[i].style.display = "none"
        continue
      }
      else if (a.attributes[0].value == "locations") {
        txtValue.toUpperCase().indexOf(filter) > -1 ? li[i].style.display = "" : li[i].style.display = "none";
        continue
      }
      else {
        li[i].style.display = "none";
        continue
      }
    }
    else if (selection == "First Album Date") {
      if (a.attributes[0].value == "groups") {
        attributesValues[1].toString().toUpperCase().indexOf(filter) > -1 ? li[i].style.display = "" : li[i].style.display = "none"
        continue
      }
      else {
        li[i].style.display = "none";
        continue
      }
    }
    else if (selection == "Creation Date") {
      if (a.attributes[0].value == "groups") {
        attributesValues[0].toString().toUpperCase().indexOf(filter) > -1 ? li[i].style.display = "" : li[i].style.display = "none"
        continue
      }
      else {
        li[i].style.display = "none";
        continue
      }
    }


    if (txtValue.toUpperCase().indexOf(filter) > -1) {
      li[i].style.display = "";
    } else {
      if (attributesValues != null) {
        var found = false

        for (var m = 0; m < 4; m++) {
          if (attributesValues[m].toString().toUpperCase().indexOf(filter) > -1) {
            li[i].style.display = "";
            found = true
            break
          }
        }

        if (found == true) continue;
      }

      li[i].style.display = "none";
    }
  }

  CheckStyle()
}