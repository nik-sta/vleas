package io.viascom.vleas.models.mavencentral

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName
import java.io.Serializable

class ResponseHeader : Serializable {

    @SerializedName("status")
    @Expose
    val status: Int? = null

    @SerializedName("QTime")
    @Expose
    val qTime: Int? = null

    @SerializedName("params")
    @Expose
    val params: Params? = null
}