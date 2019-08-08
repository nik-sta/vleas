package io.viascom.vleas.models.mavencentral

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName
import java.io.Serializable

class MavenCentralResponse : Serializable {

    @SerializedName("responseHeader")
    @Expose
    val responseHeader: ResponseHeader? = null

    @SerializedName("response")
    @Expose
    val response: Response? = null
}
