var minMembers, maxMembers;
var minAlbum, maxAlbum;
var minCreation, maxCreation;

$(document).ready(function () {
  InitFilterMembers();
  InitFilterAlbum();
  InitFilterCreation();

  $('div.tags').delegate('input:checkbox', 'change', function () {
    var counter = 0
    console.log("teste")

    $('input:checked').each(function () {
      counter = counter + 1
    });

    if (counter == 0) {
      $('.showcase-inner > article').attr("check-box-filter", "false")
    }

    var selector = $('input:checked').map(function () {
      return $(this).attr('data-concert-places');
    }).get();
    // console.log(selector);
    function filterCategories(elem) {
      var elemCats = elem.attr('data-concert-places');

      if (elemCats) {
        elemCats = elemCats.split(';');
      } else {
        elemCats = Array();
      }
      for (var i = 0; i < selector.length; i++) {
        // console.log("testing " + selector[i] + " in:");
        // console.log(elemCats);
        if (jQuery.inArray(selector[i], elemCats) != -1) {
          // console.log('false');
          $(this).attr("check-box", "test");
          return true;
        }
      }
      // console.log('true');
      return false;
    }
    $('.showcase-inner > article').each(function (i, elem) {
      if (filterCategories(jQuery(elem))) {
        var test = $(elem).attr("range-filter");

        $(elem).attr("check-box-filter", "true");
      } else {
        $(elem).attr("check-box-filter", "false");
      }
    });

    ShowByFilter();

  }).find('input:checkbox').change();
});

function ShowByFilter() {
  var rangFilterApplied = false
  var checkBoxFilterApplied = false
  $("#groups article").hide();

  $("#groups article").each(function () {
    if ($(this).attr("range-filter") == "true") {
      rangFilterApplied = true
    }

    if ($(this).attr("check-box-filter") == "true") {
      checkBoxFilterApplied = true
    }
  })

  if (!rangFilterApplied && !checkBoxFilterApplied) {
    $("#groups article").show();
  } else if (rangFilterApplied && !checkBoxFilterApplied) {
    $("#groups article").filter(function () {
      var tempVar = $(this).attr("range-filter");

      if (tempVar == "true") {
        return true
      } else {
        return false
      }
    }).show();
  } else if (!rangFilterApplied && checkBoxFilterApplied) {
    $("#groups article").filter(function () {
      var tempVar = $(this).attr("check-box-filter");

      if (tempVar == "true") {
        return true
      } else {
        return false
      }
    }).show();
  } else {
    $("#groups article").filter(function () {
      var tempVar1 = $(this).attr("check-box-filter");
      var tempVar2 = $(this).attr("range-filter");

      if (tempVar1 == "true" && tempVar2 == "true") {
        return true
      } else {
        return false
      }
    }).show();
  }
}

function Filter() {
  $("#groups article").attr("range-filter", "false").filter(function () {
    var members = parseInt($(this).data("members-counter"), 10);
    var album = parseInt($(this).data("album-creation-year"), 10);
    var creation = parseInt($(this).data("group-creation-year"), 10);
    return (members >= minMembers && members <= maxMembers) && (album >= minAlbum && album <= maxAlbum) && (creation >= minCreation && creation <= maxCreation);
  }).attr("range-filter", "true");

  ShowByFilter();
}

function InitFilterMembers() {
  var options = {
    range: true,
    min: 1,
    max: 8,
    values: [1, 8],
    slide: function (event, ui) {
      var min = ui.values[0],
        max = ui.values[1];

      $("#amount").val(min + " - " + max);

      minMembers = min
      maxMembers = max

      Filter()
    }
  }, min, max;

  $("#slider-range-amount").slider(options);
  min = $("#slider-range-amount").slider("values", 0);
  max = $("#slider-range-amount").slider("values", 1);

  $("#amount").val(min + " - " + max);
  minMembers = min
  maxMembers = max
  Filter()
};

function InitFilterAlbum() {
  var options = {
    range: true,
    min: 1963,
    max: 2018,
    values: [1963, 2018],
    slide: function (event, ui) {
      var min = ui.values[0],
        max = ui.values[1];

      $("#album").val(min + " - " + max);

      minAlbum = min
      maxAlbum = max

      Filter()
    }
  }, min, max;

  $("#slider-range-album").slider(options);

  min = $("#slider-range-album").slider("values", 0);
  max = $("#slider-range-album").slider("values", 1);

  $("#album").val(min + " - " + max);

  minAlbum = min
  maxAlbum = max
  Filter()
};

function InitFilterCreation() {
  var options = {
    range: true,
    min: 1958,
    max: 2015,
    values: [1958, 2015],
    slide: function (event, ui) {
      var min = ui.values[0],
        max = ui.values[1];

      $("#creation").val(min + " - " + max);

      minCreation = min
      maxCreation = max

      Filter()
    }
  }, min, max;

  $("#slider-range-creation").slider(options);

  min = $("#slider-range-creation").slider("values", 0);
  max = $("#slider-range-creation").slider("values", 1);

  $("#creation").val(min + " - " + max);

  minCreation = min
  maxCreation = max
  Filter()
};