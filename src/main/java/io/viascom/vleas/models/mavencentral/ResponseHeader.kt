package io.viascom.vleas.models.mavencentral

import java.io.Serializable

class ResponseHeader(
    val status: Int,
    val qTime: Int,
    val params: Params
) : Serializable