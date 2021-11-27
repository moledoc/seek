" find patterns/files/dirs/TODOs cwd recursively.
" to see the source of `seek` go to https://gitlab.com/utt_meelis/seek
function! Seek(word,flag)
    if a:word == "" && a:flag == ""
        let l:word = expand('<cWORD>')
    elseif a:word != ""
        let l:word = substitute(a:word," ","\" \"","g")
    elseif a:word == "" && a:flag == "-todo"
"         value l:word with something from the todo flag
        let l:word = "TODO:"
    endif
    execute "vert 100 new"
    execute "r ! seek -ignore=\"\\\\.git\|\\\\.seekbuf\|\\\\.swp\" " . a:flag . " \"" . l:word . "\""
    execute "w! .seekbuf"
    execute "e .seekbuf"
endfunction
" find word under cursor recursively from current working directory
nnoremap <C-f> :call Seek("","")<CR>
" find given word(s) recursively from current working directory
if !exists(":F")
    command -nargs=* F call Seek("<args>","")
endif
" find todos (+ given word(s)) recursively from current working directory
if !exists(":T")
    command -nargs=* T call Seek("<args>","-todo")
endif
