package io.viascom.vleas.models.mavencentral

import com.google.gson.annotations.Expose
import com.google.gson.annotations.SerializedName
import java.io.Serializable

class Params : Serializable {

    @SerializedName("q")
    @Expose
    val q: String? = null

    @SerializedName("core")
    @Expose
    val core: String? = null

    @SerializedName("indent")
    @Expose
    val indent: String? = null

    @SerializedName("fl")
    @Expose
    val fl: String? = null

    @SerializedName("start")
    @Expose
    val start: String? = null

    @SerializedName("sort")
    @Expose
    val sort: String? = null

    @SerializedName("rows")
    @Expose
    val rows: String? = null

    @SerializedName("wt")
    @Expose
    val wt: String? = null

    @SerializedName("version")
    @Expose
    val version: String? = null
}
