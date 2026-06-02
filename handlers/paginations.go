package handlers

import (
	"net/http"
	 "groupie-tracker/api"
	 "strconv"
)


type PaginationInfo struct{
   Artists []api.FullArtistInfo
   NextPage int
   PrevPage int
   Pages []int
   Start int
   End int
   TotalArtists int
   PageNo int
   TotalPages int
}



func Pagination(w http.ResponseWriter, r *http.Request)PaginationInfo{
	artists := api.AllArtist()
	limit := 9
    totalarists := len(artists)
	pageString := r.URL.Query().Get("page")
    pageNo, err:= strconv.Atoi(pageString)
    if err != nil || pageNo < 1 {
	    pageNo = 1
    }

	var next int
	var prev int

	
	start:= (pageNo -1) * limit

	if start >= totalarists{
	start = 0
	pageNo = 1
	}

	end := start + limit
    
	if end > totalarists{
		end = totalarists
	}

	displayArtist := artists[start:end]



	
	totalPages := totalarists / limit
    remainder := totalarists % limit

	if remainder != 0{
		totalPages += 1
	}


	pages := make([]int, totalPages)


	for i :=0; i < len(pages); i++{
		pages[i] = i+1
	}
    
    if pageNo <= 1{
       prev = 1 
	}else{
		prev = pageNo -1
	}


	if pageNo >=totalPages{
		next = totalPages
	}else{
		next = pageNo + 1
	}


	data := PaginationInfo{
		Artists : displayArtist,
		NextPage:next,
		PrevPage:prev,
		Pages:pages,
		Start:start,
		End:end,
		TotalArtists:totalarists,
		PageNo:pageNo,
        TotalPages:totalPages,
	}

  return data

}