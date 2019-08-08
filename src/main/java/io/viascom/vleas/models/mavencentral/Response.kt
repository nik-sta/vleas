package io.viascom.vleas.models.mavencentral

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName
import java.io.Serializable

class Response : Serializable {

    @SerializedName("numFound")
    @Expose
    val numFound: Int? = null

    @SerializedName("start")
    @Expose
    val start: Int? = null

    @SerializedName("docs")
    @Expose
    val docs: List<Doc>? = null
}