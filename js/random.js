function RandomGroup() {
  var max = document.getElementById("random").getAttribute("data-amount");
  document.getElementById("random").href = Math.floor(Math.random() * max + 1);
}