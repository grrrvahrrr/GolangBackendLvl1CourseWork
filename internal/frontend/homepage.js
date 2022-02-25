function shortenURL(event) {

}


function getData(event) {
    var url = "http://localhost:8000/getData";

    var xhr = new XMLHttpRequest();
    xhr.open("POST", url);

    xhr.setRequestHeader("Accept", "application/json");
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function () {
       if (xhr.readyState === 4) {
          console.log(xhr.responseText);
       }};

       document.getElementById('frm2').addEventListener('submit', (e) => {
        e.preventDefault();

        const formData = new FormData(e.target);
        const data = Array.from(formData.entries()).reduce((memo, [key, value]) => ({
          ...memo,
          [key]: value,
        }), {});
        xhr.send(JSON.stringify(data));
      });
}
