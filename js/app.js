function findArtist(art) {

    var tapedValue = document.getElementById("searchBar").value;
    var found = [];

    for (var i = 0; i < art.length; i++) {
        var name = art[i].Name.toLowerCase();
        if (name != "") {
            if (name.search(tapedValue.toLowerCase()) != -1) {
                found.push(art[i]);
            }
        }
    }

    var artTemplate1 = '';

    for (var i = 0; i < found.length; i++) {
        var members = found[i].Members;

        artTemplate1 += '<div class="box">' +
            '<img class="image" src="' + found[i].Image + '" alt="artist">' +
            '<div class="">' +
            '<h2>' + found[i].Name + '</h2>' +
            '<p>' + found[i].CreationDate + '</p>' +
            <!-- information cachée : input = hidden-->
            '<div class="membres">';

        var artTemplate2 = '';
        for (var j = 0; j < members.length; j++) {
            artTemplate2 += '<span>' + members[j] + '</span>';
        }
        var artTemplate3 =
            '</div>' +
            '</div>' +
            '<div class="infosupp">' +
            '<a href="/map?id=' + found[i].Id + '">Savoir + </a>' +
            '</div>' +
            '</div>';

        artTemplate1 += artTemplate2 + artTemplate3
    }
    var artContainer = document.getElementById('artistContainer');
    artContainer.innerHTML = artTemplate1;
}
