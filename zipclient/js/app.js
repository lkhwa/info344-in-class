// fetch("http://localhost:4000/memory")
// .then(function(response) {
//     console.log(response);
//     return response.json();
// }). then(function(data) {
    
//     console.log(data);
//     console.log(data.Alloc);
//     document.getElementById("memory").innerHTML = data.Alloc;
// })
$( document ).ready()
// function doAjax() {
//     $.ajax({
//         type: 'GET',
//         url: 'http://localhost:4000/memory',
//         dataType: 'json',
//         success: function(data) {
//             //console.log(data);
//             //console.log(data.Alloc);
//             document.getElementById("memory").innerHTML = data.Alloc;
//         },
//         complete: function (data) {
            
//             setTimeout(doAjax, 10000);
//         }

//     })
// }
// setTimeout(doAjax, 10000)


// $('form').submit(function(e) {
//     console.log($('input').val());
//     var name = $('input').val();
//     e.preventDefault();
//     // var xhttp = new XMLHttpRequest();
//     // xhttp.open("POST", "http://localhost:4000/hello?name=", true);
//     // xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
//     // xhttp.send(name);

//     fetch('http://localhost:4000/hello?name=' + name)
//     .then(function(response) {
//         console.log(response);
//         return response.text();
        
//     }) .then(function(data) {
//         console.log(data);
//         document.getElementById("name").innerHTML = data;
//     })

// })


$('form').submit(function(e) {
    e.preventDefault();
    var cityVal = $('#city').val();
    var stateVal = $('#state').val();
    var url = 'http://localhost:4000/zips/' + cityVal;
    console.log($('#city').val());
    console.log($('#state').val());
    fetch(url)
    .then(function(response) {
        console.log(response);
        return response.json();
    }).then(function(data) {
        console.log(data);
        var zipCodes = [];
        for(var stateData in data) {
            //console.log(data[stateData])
             if (stateVal.toLowerCase() == data[stateData]['State'].toLowerCase()) {
                 console.log(data[stateData]['Code']);
                 zipCodes.push(data[stateData]['Code']);
             }
        }
        console.log(zipCodes);
        document.getElementById("zipcode").innerHTML = zipCodes;
    }) 
    
})



