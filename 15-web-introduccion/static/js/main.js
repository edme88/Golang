console.log("Hola, desde main.js");

let imgArray = ["hand_rock", "hand_paper", "hand_scissors"];

function choose(x) {
  fetch(`/play?c=${x}`)
    .then((reponse) => reponse.json())
    .then((data) => {
      console.log(data);

      if (x === 0) {
        document.getElementById("player_choice").innerHTML =
          "El jugador eligió PIEDRA.";
        document.getElementById("img_computer");
      } else if (x === 1) {
        document.getElementById("player_choice").innerHTML =
          "El jugador eligió PAPEL.";
      } else {
        document.getElementById("player_choice").innerHTML =
          "El jugador eligió TIJERA.";
      }

      document.getElementById("player_score").innerHTML = data.player_score;
      document.getElementById("computer_score").innerHTML = data.computer_score;

      document.getElementById("computer_choice").innerHTML =
        data.computer_choice;
      document.getElementById("round_result").innerHTML = data.round_result;
      document.getElementById("round_message").innerHTML = data.message;

      let imgElement = document.getElementById("img_computer");
      imgElement.src = `static/img/${imgArray[data.computer_choice_int]}.png`;
    });
}
