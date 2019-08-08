package io.viascom.vleas.models.mavencentral

import java.io.Serializable

class Response(
    val numFound: Int,
    val start: Int,
    val docs: List<Doc>
) : Serializable