$(function(){
    var el = jQuery('input').autocomplete({
      minLength: 2,
      source: function( request, response ) {
        jQuery.ajax({
            url: 'data.json',
            dataType: "json",
            data: {
                    term: request.term
            },
            success: function( data ) {
              response( jQuery.map( json.aa, function( item ) {
                return {
                  //id: item.id,
                  value: item.name,
                  label: item.name,
                  company: item.company
                }
            }));
          }
        }).error(function(){console.log('err', arguments)});
      },
      select: function( event, ui ) {
            return false;
      }
    })
  });