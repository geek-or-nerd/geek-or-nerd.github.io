$(function() {
  return $('.ta').typeahead({
    name: 'names',
    prefetch: 'data/terms.json',
    limit: 10
  });
});