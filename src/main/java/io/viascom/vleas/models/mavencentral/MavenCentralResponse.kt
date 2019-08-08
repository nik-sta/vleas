package io.viascom.vleas.models.mavencentral

import java.io.Serializable

class MavenCentralResponse(
    val responseHeader: ResponseHeader,
    val response: Response
) : Serializable
