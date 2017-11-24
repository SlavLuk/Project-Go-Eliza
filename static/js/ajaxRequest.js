
    $(function () {

      var scroll = 200;

      $("#user-input-form").submit(
        function (event) {

          event.preventDefault();

          $.get('/user-input', {

            value: $('#user-input').val().trim()

          })
            .done(function (data) {

              var userData = $('#user-input').val().trim();

              userData = "You: " + userData;

              $('#user-input').val("");

              $('#output-area').append("<li class='bg-info text-white rounded float-left'>" + userData + "</li>");

              setTimeout(function () {

                $('#output-area').append("<li class='bg-success text-white rounded float-right'>" + data + "</li>");

                scroll += 200;
                window.scrollTo(0, scroll);

              }, 1000);
            })
            .fail(function () {
              alert("error");

            });

        });
    });