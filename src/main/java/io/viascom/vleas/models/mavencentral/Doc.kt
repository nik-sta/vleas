package io.viascom.vleas.models.mavencentral

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName
import java.io.Serializable

class Doc : Serializable {

    @SerializedName("id")
    @Expose
    val id: String? = null

    @SerializedName("g")
    @Expose
    val g: String? = null

    @SerializedName("a")
    @Expose
    val a: String? = null

    @SerializedName("v")
    @Expose
    val v: String? = null

    @SerializedName("p")
    @Expose
    val p: String? = null

    @SerializedName("timestamp")
    @Expose
    val timestamp: String? = null

    @SerializedName("ec")
    @Expose
    val ec: List<String>? = null

    @SerializedName("tags")
    @Expose
    val tags: List<String>? = null
}