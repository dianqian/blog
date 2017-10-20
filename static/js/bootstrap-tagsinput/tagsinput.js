var citynames = new Bloodhound({
    datumTokenizer: Bloodhound.tokenizers.obj.whitespace('name'),
    queryTokenizer: Bloodhound.tokenizers.whitespace,
    prefetch: {
        url: '/admin/tags/list',
        filter: function(list) {
            return $.map(list, function(cityname) {
                return { name: cityname }; });
        }
    }
});
citynames.initialize();

/**
 * Typeahead
 */
var elt = $('.tags_typeahead > > input');
elt.tagsinput({
    typeaheadjs: {
        name: 'citynames',
        displayKey: 'name',
        valueKey: 'name',
        source: citynames.ttAdapter()
    }
});